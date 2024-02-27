package main

import (
	"ManeBackend/internal/env"
	"ManeBackend/internal/jwt"
	"ManeBackend/models"
	"ManeBackend/pb"
	"ManeBackend/service"
	"context"
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type HealthCheck struct {
	pb.UnimplementedHealthCheckServer
}

func (s *HealthCheck) Check(_ context.Context, request *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("Recevied healthcheck request %v", request.GetService())
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthCheck) Watch(_ *pb.HealthCheckRequest, stream pb.HealthCheck_WatchServer) error {
	err := stream.Send(&pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	})
	return err
}

func main() {
	config := env.GetConfig()
	err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	jwtManager := jwt.NewJWTManager(config.JWT_SECRET)

	launch := launcher.New().Leakless(false).Bin("/usr/bin/chromium").Set("no-sandbox", "true").MustLaunch()
	//launch := launcher.New().Headless(false).Devtools(true).Leakless(false).Bin("/opt/homebrew/bin/chromium").MustLaunch()
	browser := rod.New().ControlURL(launch).MustConnect()
	log.Print("Launched and connected to browser")
	defer browser.MustClose()
	router := browser.HijackRequests()
	defer router.Stop()
	router.MustAdd("*", func(ctx *rod.Hijack) {
		switch resourceType := ctx.Request.Type(); resourceType {
		case proto.NetworkResourceTypeStylesheet, proto.NetworkResourceTypeImage, proto.NetworkResourceTypeFont, proto.NetworkResourceTypeMedia:
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		default:
			ctx.ContinueRequest(&proto.FetchContinueRequest{})
			return
		}
	})
	go router.Run()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.PORT))
	if err != nil {
		log.Fatal(err)
	}

	mainService := service.NewMainService(jwtManager)
	initService := service.NewInitService(browser)
	interceptor := service.NewAuthInterceptor(jwtManager)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()))

	pb.RegisterMainServiceServer(s, mainService)
	pb.RegisterHealthCheckServer(s, &HealthCheck{})
	pb.RegisterInitServiceServer(s, initService)

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Server started on port %v", config.PORT)
	defer models.Close()
	defer s.Stop()
}

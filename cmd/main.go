package main

import (
	"ManeBackend/internal/env"
	"ManeBackend/internal/jwt"
	"ManeBackend/models"
	"ManeBackend/pb"
	"ManeBackend/service"
	"context"
	"fmt"
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

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", config.PORT))
	if err != nil {
		log.Fatal(err)
	}

	mainService := service.NewMainService(jwtManager)
	interceptor := service.NewAuthInterceptor(jwtManager)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()))

	pb.RegisterMainServiceServer(s, mainService)
	pb.RegisterHealthCheckServer(s, &HealthCheck{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Printf("Server started on port %v", config.PORT)
	defer models.Close()
	defer s.Stop()
}

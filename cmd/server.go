package main

import (
	"ManeBackend/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"os"
)

type Service struct {
	pb.UnimplementedServiceServer
}

func (s *Service) GetUpdatedURLs(ctx context.Context, request *pb.GetUpdatedURLsRequest) (*pb.GetUpdatedURLsResponse, error) {
	if request.GetVersionTimestamp() == nil {
		return &pb.GetUpdatedURLsResponse{
			LatestURLs: nil,
		}, nil
	}
	println(request.GetVersionTimestamp().AsTime().String())
	ts := request.GetVersionTimestamp().AsTime()
	println(ts.String())
	URLs := make(map[string]string)
	URLs["home"] = "12312f.com"
	LatestUrls := &pb.URLsList{
		VersionTimestamp: timestamppb.Now(),
		URLs:             URLs,
	}
	return &pb.GetUpdatedURLsResponse{
		LatestURLs: LatestUrls,
	}, nil
}

type HealthCheck struct {
	pb.UnimplementedHealthServer
}

func (s *HealthCheck) Check(_ context.Context, request *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("Recevied healthcheck request %v", request.GetService())
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthCheck) Watch(request *pb.HealthCheckRequest, stream pb.Health_WatchServer) error {
	err := stream.Send(&pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	})
	return err
}

func main() {
	PORT, exists := os.LookupEnv("PORT")
	if !exists {
		PORT = "8080"
	}
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &Service{})
	pb.RegisterHealthServer(s, &HealthCheck{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

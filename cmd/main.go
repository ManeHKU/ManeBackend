package main

import (
	"ManeBackend/pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

type Service struct {
	pb.UnimplementedServiceServer
}

type HealthCheck struct {
	pb.UnimplementedHealthServer
}

func main() {
	PORT, exists := os.LookupEnv("PORT")
	if !exists {
		PORT = "8080"
	}
	log.Printf("Read PORT env, will be listening at %v", PORT)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatal(err)
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

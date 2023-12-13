package main

import (
	"ManeBackend/pb"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

func (s *HealthCheck) Check(_ context.Context, request *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Printf("Recevied healthcheck request %v", request.GetService())
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthCheck) Watch(_ *pb.HealthCheckRequest, stream pb.Health_WatchServer) error {
	err := stream.Send(&pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	})
	return err
}

func (s *Service) GetUpdatedURLs(_ context.Context, request *pb.GetUpdatedURLsRequest) (*pb.GetUpdatedURLsResponse, error) {
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
	LatestUrls := &pb.GetUpdatedURLsResponse_URLsList{
		VersionTimestamp: timestamppb.Now(),
		URLs:             URLs,
	}
	return &pb.GetUpdatedURLsResponse{
		LatestURLs: LatestUrls,
	}, nil
}

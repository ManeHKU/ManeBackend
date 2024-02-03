package service

import (
	"ManeBackend/constants"
	"ManeBackend/internal/jwt"
	"ManeBackend/models"
	"ManeBackend/pb"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type MainService struct {
	jwtManager *jwt.Manager
	pb.UnimplementedMainServiceServer
}

func (s *MainService) GetUpdatedURLs(_ context.Context, request *pb.GetUpdatedURLsRequest) (*pb.GetUpdatedURLsResponse, error) {
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

func (s *MainService) UpdateUserInfo(context context.Context, request *pb.UpdateUserInfoRequest) (*emptypb.Empty, error) {
	if request.GetUid() == 0 && request.GetFullName() == "" {
		log.Printf("empty request not doing anything")
		return nil, status.Errorf(codes.InvalidArgument, "empty request")
	}
	jwtClaims := context.Value(constants.JWTClaimsKey).(*jwt.UserClaims)
	userUUID := jwtClaims.Subject
	if userUUID == "" {
		log.Printf("empty uuid or error")
		return &emptypb.Empty{}, nil
	}

	log.Printf("current user id: %v", userUUID)
	if !models.CheckUserExist(userUUID) {
		log.Printf("user not exist")
		return nil, status.Errorf(codes.NotFound, "user not exist")
	}
	log.Printf("user %v exists", userUUID)

	if err := models.UpdateUserInfo(userUUID, request.GetUid(), request.GetFullName()); err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func NewMainService(jwtManager *jwt.Manager) *MainService {
	return &MainService{jwtManager: jwtManager}
}

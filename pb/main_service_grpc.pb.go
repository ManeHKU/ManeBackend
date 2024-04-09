// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: main_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MainService_GetUpdatedURLs_FullMethodName       = "/service.MainService/GetUpdatedURLs"
	MainService_UpdateUserInfo_FullMethodName       = "/service.MainService/UpdateUserInfo"
	MainService_UpsertTakenCourses_FullMethodName   = "/service.MainService/UpsertTakenCourses"
	MainService_ListCourses_FullMethodName          = "/service.MainService/ListCourses"
	MainService_SearchCourses_FullMethodName        = "/service.MainService/SearchCourses"
	MainService_GetCourseDetails_FullMethodName     = "/service.MainService/GetCourseDetails"
	MainService_AddReview_FullMethodName            = "/service.MainService/AddReview"
	MainService_ListLatestEvents_FullMethodName     = "/service.MainService/ListLatestEvents"
	MainService_AddEvent_FullMethodName             = "/service.MainService/AddEvent"
	MainService_ApplyEvent_FullMethodName           = "/service.MainService/ApplyEvent"
	MainService_ListUserAppliedEvent_FullMethodName = "/service.MainService/ListUserAppliedEvent"
)

// MainServiceClient is the client API for MainService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainServiceClient interface {
	GetUpdatedURLs(ctx context.Context, in *GetUpdatedURLsRequest, opts ...grpc.CallOption) (*GetUpdatedURLsResponse, error)
	UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpsertTakenCourses(ctx context.Context, in *UpsertTakenCoursesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (*CoursesResponse, error)
	SearchCourses(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*CoursesResponse, error)
	GetCourseDetails(ctx context.Context, in *GetCourseDetailRequest, opts ...grpc.CallOption) (*GetCourseDetailResponse, error)
	AddReview(ctx context.Context, in *AddReviewRequest, opts ...grpc.CallOption) (*AddReviewResponse, error)
	ListLatestEvents(ctx context.Context, in *ListLatestEventsRequest, opts ...grpc.CallOption) (*ListLatestEventsResponse, error)
	AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error)
	ApplyEvent(ctx context.Context, in *ApplyEventRequest, opts ...grpc.CallOption) (*ApplyEventResponse, error)
	ListUserAppliedEvent(ctx context.Context, in *ListUserAppliedEventRequest, opts ...grpc.CallOption) (*ListUserAppliedEventResponse, error)
}

type mainServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMainServiceClient(cc grpc.ClientConnInterface) MainServiceClient {
	return &mainServiceClient{cc}
}

func (c *mainServiceClient) GetUpdatedURLs(ctx context.Context, in *GetUpdatedURLsRequest, opts ...grpc.CallOption) (*GetUpdatedURLsResponse, error) {
	out := new(GetUpdatedURLsResponse)
	err := c.cc.Invoke(ctx, MainService_GetUpdatedURLs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) UpdateUserInfo(ctx context.Context, in *UpdateUserInfoRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MainService_UpdateUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) UpsertTakenCourses(ctx context.Context, in *UpsertTakenCoursesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MainService_UpsertTakenCourses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) ListCourses(ctx context.Context, in *ListCoursesRequest, opts ...grpc.CallOption) (*CoursesResponse, error) {
	out := new(CoursesResponse)
	err := c.cc.Invoke(ctx, MainService_ListCourses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) SearchCourses(ctx context.Context, in *SearchCourseRequest, opts ...grpc.CallOption) (*CoursesResponse, error) {
	out := new(CoursesResponse)
	err := c.cc.Invoke(ctx, MainService_SearchCourses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) GetCourseDetails(ctx context.Context, in *GetCourseDetailRequest, opts ...grpc.CallOption) (*GetCourseDetailResponse, error) {
	out := new(GetCourseDetailResponse)
	err := c.cc.Invoke(ctx, MainService_GetCourseDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) AddReview(ctx context.Context, in *AddReviewRequest, opts ...grpc.CallOption) (*AddReviewResponse, error) {
	out := new(AddReviewResponse)
	err := c.cc.Invoke(ctx, MainService_AddReview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) ListLatestEvents(ctx context.Context, in *ListLatestEventsRequest, opts ...grpc.CallOption) (*ListLatestEventsResponse, error) {
	out := new(ListLatestEventsResponse)
	err := c.cc.Invoke(ctx, MainService_ListLatestEvents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) AddEvent(ctx context.Context, in *AddEventRequest, opts ...grpc.CallOption) (*AddEventResponse, error) {
	out := new(AddEventResponse)
	err := c.cc.Invoke(ctx, MainService_AddEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) ApplyEvent(ctx context.Context, in *ApplyEventRequest, opts ...grpc.CallOption) (*ApplyEventResponse, error) {
	out := new(ApplyEventResponse)
	err := c.cc.Invoke(ctx, MainService_ApplyEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainServiceClient) ListUserAppliedEvent(ctx context.Context, in *ListUserAppliedEventRequest, opts ...grpc.CallOption) (*ListUserAppliedEventResponse, error) {
	out := new(ListUserAppliedEventResponse)
	err := c.cc.Invoke(ctx, MainService_ListUserAppliedEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServiceServer is the server API for MainService service.
// All implementations must embed UnimplementedMainServiceServer
// for forward compatibility
type MainServiceServer interface {
	GetUpdatedURLs(context.Context, *GetUpdatedURLsRequest) (*GetUpdatedURLsResponse, error)
	UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*emptypb.Empty, error)
	UpsertTakenCourses(context.Context, *UpsertTakenCoursesRequest) (*emptypb.Empty, error)
	ListCourses(context.Context, *ListCoursesRequest) (*CoursesResponse, error)
	SearchCourses(context.Context, *SearchCourseRequest) (*CoursesResponse, error)
	GetCourseDetails(context.Context, *GetCourseDetailRequest) (*GetCourseDetailResponse, error)
	AddReview(context.Context, *AddReviewRequest) (*AddReviewResponse, error)
	ListLatestEvents(context.Context, *ListLatestEventsRequest) (*ListLatestEventsResponse, error)
	AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error)
	ApplyEvent(context.Context, *ApplyEventRequest) (*ApplyEventResponse, error)
	ListUserAppliedEvent(context.Context, *ListUserAppliedEventRequest) (*ListUserAppliedEventResponse, error)
	mustEmbedUnimplementedMainServiceServer()
}

// UnimplementedMainServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMainServiceServer struct {
}

func (UnimplementedMainServiceServer) GetUpdatedURLs(context.Context, *GetUpdatedURLsRequest) (*GetUpdatedURLsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUpdatedURLs not implemented")
}
func (UnimplementedMainServiceServer) UpdateUserInfo(context.Context, *UpdateUserInfoRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserInfo not implemented")
}
func (UnimplementedMainServiceServer) UpsertTakenCourses(context.Context, *UpsertTakenCoursesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertTakenCourses not implemented")
}
func (UnimplementedMainServiceServer) ListCourses(context.Context, *ListCoursesRequest) (*CoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourses not implemented")
}
func (UnimplementedMainServiceServer) SearchCourses(context.Context, *SearchCourseRequest) (*CoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCourses not implemented")
}
func (UnimplementedMainServiceServer) GetCourseDetails(context.Context, *GetCourseDetailRequest) (*GetCourseDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourseDetails not implemented")
}
func (UnimplementedMainServiceServer) AddReview(context.Context, *AddReviewRequest) (*AddReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReview not implemented")
}
func (UnimplementedMainServiceServer) ListLatestEvents(context.Context, *ListLatestEventsRequest) (*ListLatestEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLatestEvents not implemented")
}
func (UnimplementedMainServiceServer) AddEvent(context.Context, *AddEventRequest) (*AddEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}
func (UnimplementedMainServiceServer) ApplyEvent(context.Context, *ApplyEventRequest) (*ApplyEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyEvent not implemented")
}
func (UnimplementedMainServiceServer) ListUserAppliedEvent(context.Context, *ListUserAppliedEventRequest) (*ListUserAppliedEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserAppliedEvent not implemented")
}
func (UnimplementedMainServiceServer) mustEmbedUnimplementedMainServiceServer() {}

// UnsafeMainServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServiceServer will
// result in compilation errors.
type UnsafeMainServiceServer interface {
	mustEmbedUnimplementedMainServiceServer()
}

func RegisterMainServiceServer(s grpc.ServiceRegistrar, srv MainServiceServer) {
	s.RegisterService(&MainService_ServiceDesc, srv)
}

func _MainService_GetUpdatedURLs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUpdatedURLsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).GetUpdatedURLs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_GetUpdatedURLs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).GetUpdatedURLs(ctx, req.(*GetUpdatedURLsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_UpdateUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).UpdateUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_UpdateUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).UpdateUserInfo(ctx, req.(*UpdateUserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_UpsertTakenCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertTakenCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).UpsertTakenCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_UpsertTakenCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).UpsertTakenCourses(ctx, req.(*UpsertTakenCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_ListCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).ListCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_ListCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).ListCourses(ctx, req.(*ListCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_SearchCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).SearchCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_SearchCourses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).SearchCourses(ctx, req.(*SearchCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_GetCourseDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourseDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).GetCourseDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_GetCourseDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).GetCourseDetails(ctx, req.(*GetCourseDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_AddReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).AddReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_AddReview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).AddReview(ctx, req.(*AddReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_ListLatestEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLatestEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).ListLatestEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_ListLatestEvents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).ListLatestEvents(ctx, req.(*ListLatestEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_AddEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).AddEvent(ctx, req.(*AddEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_ApplyEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).ApplyEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_ApplyEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).ApplyEvent(ctx, req.(*ApplyEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MainService_ListUserAppliedEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserAppliedEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServiceServer).ListUserAppliedEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MainService_ListUserAppliedEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServiceServer).ListUserAppliedEvent(ctx, req.(*ListUserAppliedEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MainService_ServiceDesc is the grpc.ServiceDesc for MainService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MainService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.MainService",
	HandlerType: (*MainServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUpdatedURLs",
			Handler:    _MainService_GetUpdatedURLs_Handler,
		},
		{
			MethodName: "UpdateUserInfo",
			Handler:    _MainService_UpdateUserInfo_Handler,
		},
		{
			MethodName: "UpsertTakenCourses",
			Handler:    _MainService_UpsertTakenCourses_Handler,
		},
		{
			MethodName: "ListCourses",
			Handler:    _MainService_ListCourses_Handler,
		},
		{
			MethodName: "SearchCourses",
			Handler:    _MainService_SearchCourses_Handler,
		},
		{
			MethodName: "GetCourseDetails",
			Handler:    _MainService_GetCourseDetails_Handler,
		},
		{
			MethodName: "AddReview",
			Handler:    _MainService_AddReview_Handler,
		},
		{
			MethodName: "ListLatestEvents",
			Handler:    _MainService_ListLatestEvents_Handler,
		},
		{
			MethodName: "AddEvent",
			Handler:    _MainService_AddEvent_Handler,
		},
		{
			MethodName: "ApplyEvent",
			Handler:    _MainService_ApplyEvent_Handler,
		},
		{
			MethodName: "ListUserAppliedEvent",
			Handler:    _MainService_ListUserAppliedEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main_service.proto",
}

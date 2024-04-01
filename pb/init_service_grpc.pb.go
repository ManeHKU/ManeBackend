// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: init_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	InitService_GetInitialConfig_FullMethodName = "/init.InitService/GetInitialConfig"
	InitService_GetSISTicket_FullMethodName     = "/init.InitService/GetSISTicket"
)

// InitServiceClient is the client API for InitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InitServiceClient interface {
	GetInitialConfig(ctx context.Context, in *GetInitialConfigRequest, opts ...grpc.CallOption) (*GetInitialConfigResponse, error)
	GetSISTicket(ctx context.Context, in *UserSignInRequest, opts ...grpc.CallOption) (*UserSignInResponse, error)
}

type initServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInitServiceClient(cc grpc.ClientConnInterface) InitServiceClient {
	return &initServiceClient{cc}
}

func (c *initServiceClient) GetInitialConfig(ctx context.Context, in *GetInitialConfigRequest, opts ...grpc.CallOption) (*GetInitialConfigResponse, error) {
	out := new(GetInitialConfigResponse)
	err := c.cc.Invoke(ctx, InitService_GetInitialConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *initServiceClient) GetSISTicket(ctx context.Context, in *UserSignInRequest, opts ...grpc.CallOption) (*UserSignInResponse, error) {
	out := new(UserSignInResponse)
	err := c.cc.Invoke(ctx, InitService_GetSISTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InitServiceServer is the server API for InitService service.
// All implementations must embed UnimplementedInitServiceServer
// for forward compatibility
type InitServiceServer interface {
	GetInitialConfig(context.Context, *GetInitialConfigRequest) (*GetInitialConfigResponse, error)
	GetSISTicket(context.Context, *UserSignInRequest) (*UserSignInResponse, error)
	mustEmbedUnimplementedInitServiceServer()
}

// UnimplementedInitServiceServer must be embedded to have forward compatible implementations.
type UnimplementedInitServiceServer struct {
}

func (UnimplementedInitServiceServer) GetInitialConfig(context.Context, *GetInitialConfigRequest) (*GetInitialConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInitialConfig not implemented")
}
func (UnimplementedInitServiceServer) GetSISTicket(context.Context, *UserSignInRequest) (*UserSignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSISTicket not implemented")
}
func (UnimplementedInitServiceServer) mustEmbedUnimplementedInitServiceServer() {}

// UnsafeInitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InitServiceServer will
// result in compilation errors.
type UnsafeInitServiceServer interface {
	mustEmbedUnimplementedInitServiceServer()
}

func RegisterInitServiceServer(s grpc.ServiceRegistrar, srv InitServiceServer) {
	s.RegisterService(&InitService_ServiceDesc, srv)
}

func _InitService_GetInitialConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInitialConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitServiceServer).GetInitialConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InitService_GetInitialConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitServiceServer).GetInitialConfig(ctx, req.(*GetInitialConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InitService_GetSISTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserSignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InitServiceServer).GetSISTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InitService_GetSISTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InitServiceServer).GetSISTicket(ctx, req.(*UserSignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InitService_ServiceDesc is the grpc.ServiceDesc for InitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "init.InitService",
	HandlerType: (*InitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInitialConfig",
			Handler:    _InitService_GetInitialConfig_Handler,
		},
		{
			MethodName: "GetSISTicket",
			Handler:    _InitService_GetSISTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "init_service.proto",
}

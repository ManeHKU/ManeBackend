package service

import (
	"ManeBackend/constants"
	"ManeBackend/internal/jwt"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"strings"
)

type AuthInterceptor struct {
	jwtManager *jwt.Manager
}

func NewAuthInterceptor(jwtManager *jwt.Manager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		newCtx, err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(newCtx, req)
	}
}

func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		newCtx, err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}
		newStream := &grpc_middleware.WrappedServerStream{
			ServerStream:   stream,
			WrappedContext: newCtx,
		}

		return handler(srv, newStream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) (context.Context, error) {
	// only main services requires authorization
	if !strings.HasPrefix(method, "/service.MainService/") {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return ctx, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	_, accessToken, found := strings.Cut(values[0], "Bearer ")

	if !found {
		return ctx, status.Errorf(codes.Unauthenticated, "invalid token format")
	}

	claims, err := interceptor.jwtManager.VerifyUserToken(accessToken)
	if err != nil {
		log.Printf("access token is invalid: %v", err)
		return ctx, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if strings.EqualFold(claims.Role, "authenticated") {
		ctx = context.WithValue(ctx, constants.JWTClaimsKey, claims)
		return ctx, nil
	}

	return ctx, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}

package grpc

import (
	baseGRPC "cms/grpc/module/base/delivery/grpc"
	"cms/grpc/module/user/usecase"
	pb "cms/protocgen/core/v1/auth"

	"google.golang.org/grpc"
)

type AuthService struct {
	authUsecase usecase.AuthUsecase
	pb.UnimplementedAuthServiceServer
	baseGRPC.GRPCHandler
}

func NewAuthService(grpcServer *grpc.Server, usecase usecase.AuthUsecase) {
	authGrpc := &AuthService{authUsecase: usecase}
	pb.RegisterAuthServiceServer(grpcServer, authGrpc)
}

package grpc

import (
	"cms/grpc/module/auth/usecase"
	baseGRPC "cms/grpc/module/base/delivery/grpc"
	pb "cms/protocgen/core/v1/auth"
	"context"

	"cms/grpc/model"

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

func (srv *AuthService) RegisterUser(
	ctx context.Context, req *pb.RegisterRequest,
) (*pb.UserResponse, error) {
	var (
		response pb.UserResponse
	)

	result, err := srv.authUsecase.RegisterUser(ctx, &model.CreateUserReq{
		BaseUser: &model.BaseUser{
			Email:    req.GetEmail(),
			Username: req.GetUsername(),
			Password: req.GetPassword(),
		},
	})

	if err != nil {
		return nil, err
	}

	if err := srv.Transform(result, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *AuthService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {

	var (
		checkVerifyToken model.VerifyCredential
		checkResponse    pb.UserResponse
	)

	if err := srv.GRPCHandler.Transform(req, &checkVerifyToken); err != nil {
		return nil, err
	}

	result, err := srv.authUsecase.VerifyCredential(ctx, checkVerifyToken)

	if err != nil {
		return nil, err
	}

	if err := srv.GRPCHandler.Transform(result, &checkResponse); err != nil {
		return nil, err
	}

	return &checkResponse, nil
}

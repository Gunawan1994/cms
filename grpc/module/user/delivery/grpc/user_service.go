package grpc

import (
	baseGRPC "cms/grpc/module/base/delivery/grpc"
	"cms/grpc/module/user/usecase"
	pb "cms/protocgen/core/v1/user"
	"context"

	"cms/grpc/model"

	"google.golang.org/grpc"
)

type UserService struct {
	UserUsecase usecase.UserUsecase
	pb.UnimplementedUserServiceServer
	baseGRPC.GRPCHandler
}

func NewUserService(grpcServer *grpc.Server, usecase usecase.UserUsecase) {
	authGrpc := &UserService{UserUsecase: usecase}
	pb.RegisterUserServiceServer(grpcServer, authGrpc)
}

func (srv *UserService) CreateUser(
	ctx context.Context, req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {
	var (
		request  model.CreateUserReq
		response pb.CreateUserResponse
	)
	if err := srv.Transform(req.User, &request.BaseUser); err != nil {
		return nil, err
	}

	result, err := srv.UserUsecase.Create(ctx, &request)

	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("User successfully created")
	response.User = &pb.User{}

	if err := srv.Transform(result.User, response.User); err != nil {
		return nil, err
	}

	return &response, nil
}

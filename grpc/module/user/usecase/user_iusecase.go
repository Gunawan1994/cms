package usecase

import (
	"cms/grpc/model"
	"context"
)

type UserUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error)
}

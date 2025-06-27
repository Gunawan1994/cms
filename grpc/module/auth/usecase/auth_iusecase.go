package usecase

import (
	"context"

	"cms/grpc/domain"
	"cms/grpc/model"
)

type AuthUsecase interface {
	// VerifyPassword(ctx context.Context, id, password string) (bool, error)
	// AuthCheckExist(ctx context.Context, req model.AuthCheckExist) (*domain.Users, error)
	VerifyCredential(ctx context.Context, req model.VerifyCredential) (*domain.User, error)
	RegisterUser(ctx context.Context, req *model.CreateUserReq) (*domain.User, error)
}

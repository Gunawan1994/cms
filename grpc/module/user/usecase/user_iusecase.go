package usecase

import (
	"context"
)

type AuthUsecase interface {
	GetUserByEmail(ctx context.Context, email string) (bool, error)
}

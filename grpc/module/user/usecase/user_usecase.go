package usecase

import (
	"context"

	"cms/grpc/model"
	"cms/grpc/module/user/repository/postgres"

	"gorm.io/gorm"
)

type UserUseCaseImpl struct {
	db   *gorm.DB
	repo postgres.UserRepository
}

func NewAuthUseCase(
	db *gorm.DB, repo postgres.UserRepository,
) AuthUsecase {
	return &UserUseCaseImpl{
		db:   db,
		repo: repo,
	}
}

func (s *UserUseCaseImpl) Create(
	ctx context.Context, req *model.CreateUserReq,
) (*model.CreateUserRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.CreateUserRes{User: body}, nil
}

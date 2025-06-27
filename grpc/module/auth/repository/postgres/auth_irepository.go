package postgres

import (
	"cms/grpc/model"
	"context"

	"cms/grpc/domain"

	"gorm.io/gorm"
)

type AuthIRepository interface {
	// repository.BaseRepository[domain.User]
	VerifyCredential(ctx context.Context, tx *gorm.DB, req model.VerifyCredential) (*domain.User, error)
}

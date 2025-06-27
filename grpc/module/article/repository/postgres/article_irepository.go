package postgres

import (
	"cms/grpc/domain"
	"context"

	"cms/grpc/model"

	"cms/grpc/module/base/repository"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	// Example operations
	repository.BaseRepository[domain.Article]
	ArticleList(
		ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
		keyword model.KeywordParam,
	) (*model.PaginationData[domain.Article], error)
	FindArticle(
		ctx context.Context, tx *gorm.DB, id int64,
	) (*domain.Article, error)
}

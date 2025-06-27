package postgres

import (
	"context"
	"strconv"

	"cms/grpc/domain"
	"cms/grpc/model"
	"cms/grpc/module/base/repository"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleRepo struct {
	repository.BaseRepository[domain.Article]
}

func NewArticleRepository() ArticleRepository {
	keywordField := []string{
		"title",
	}

	repo := repository.NewBaseRepositoryImpl[domain.Article](keywordField)
	return &ArticleRepo{
		BaseRepository: repo,
	}
}

func (r *ArticleRepo) ArticleList(
	ctx context.Context, tx *gorm.DB, page model.PaginationParam, order model.OrderParam, filter model.FilterParams,
	keyword model.KeywordParam,
) (*model.PaginationData[domain.Article], error) {
	return r.Find(ctx, tx.Preload(clause.Associations), page, order, filter, keyword)
}

func (r *ArticleRepo) FindArticle(
	ctx context.Context, tx *gorm.DB, id int64,
) (*domain.Article, error) {
	query := tx.WithContext(ctx).Preload(clause.Associations)

	result, err := r.FindByID(ctx, query, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}

	return result, nil
}

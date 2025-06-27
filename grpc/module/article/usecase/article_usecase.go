package usecase

import (
	"context"
	"errors"
	"strconv"

	"cms/grpc/model"
	"cms/grpc/module/article/repository/postgres"

	"gorm.io/gorm"
)

type ArticleUseCaseImpl struct {
	db   *gorm.DB
	repo postgres.ArticleRepository
}

func NewArticleUseCase(
	db *gorm.DB, repo postgres.ArticleRepository,
) ArticleUseCase {
	return &ArticleUseCaseImpl{
		db:   db,
		repo: repo,
	}
}

func (s *ArticleUseCaseImpl) Create(
	ctx context.Context, req *model.CreateArticleReq,
) (*model.CreateArticleRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	if err := s.repo.CreateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.CreateArticleRes{Article: body}, nil
}

func (s *ArticleUseCaseImpl) GetById(
	ctx context.Context, req *model.GetIdArticleReq,
) (*model.GetIdArticleRes, error) {
	result, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("article not found")
	}

	return &model.GetIdArticleRes{Article: result}, nil
}

func (s *ArticleUseCaseImpl) Update(
	ctx context.Context, req *model.UpdateArticleReq,
) (*model.UpdateArticleRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	body := req.ToDomain()

	data, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	body.Id = data.Id

	if err := s.repo.UpdateTx(ctx, tx, body); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.UpdateArticleRes{Article: body}, nil

}

func (s *ArticleUseCaseImpl) GetList(
	ctx context.Context, req *model.GetListArticleReq,
) (*model.GetListArticleRes, error) {
	result, err := s.repo.ArticleList(ctx, s.db, req.Page, req.Order, req.Filter, req.Keyword)
	if err != nil {
		return nil, err
	}

	return &model.GetListArticleRes{
		Data:       result.Data,
		Pagination: result.ToPagination(),
	}, nil
}

func (s *ArticleUseCaseImpl) Delete(
	ctx context.Context, req *model.DeleteArticleReq,
) (*model.DeleteArticleRes, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	data, err := s.repo.FindByID(ctx, s.db, strconv.FormatInt(req.Id, 10))
	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, errors.New("Article not found")
	}

	if err := s.repo.DeleteByIDTx(ctx, tx, strconv.FormatInt(req.Id, 10)); err != nil {
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &model.DeleteArticleRes{Article: data}, nil
}

package usecase

import (
	"context"

	"cms/grpc/model"
)

type ArticleUseCase interface {
	Create(ctx context.Context, req *model.CreateArticleReq) (*model.CreateArticleRes, error)
	GetById(ctx context.Context, req *model.GetIdArticleReq) (*model.GetIdArticleRes, error)
	GetList(ctx context.Context, req *model.GetListArticleReq) (
		*model.GetListArticleRes, error,
	)
	Update(ctx context.Context, req *model.UpdateArticleReq) (*model.UpdateArticleRes, error)
	Delete(ctx context.Context, req *model.DeleteArticleReq) (*model.DeleteArticleRes, error)
}

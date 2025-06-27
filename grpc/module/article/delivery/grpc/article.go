package grpc

import (
	"context"
	"reflect"

	"cms/grpc/model"
	"cms/grpc/module/article/usecase"
	baseGRPC "cms/grpc/module/base/delivery/grpc"
	pb "cms/protocgen/core/v1/article"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArticleService struct {
	ArticleUseCase usecase.ArticleUseCase
	pb.UnimplementedArticleServiceServer
	baseGRPC.GRPCHandler
}

func NewArticleService(grpcServer *grpc.Server, usecase usecase.ArticleUseCase) {
	attributeGrpc := &ArticleService{
		ArticleUseCase: usecase,
	}
	pb.RegisterArticleServiceServer(grpcServer, attributeGrpc)
}

func (srv *ArticleService) CreateArticle(
	ctx context.Context, req *pb.CreateArticleRequest,
) (*pb.CreateArticleResponse, error) {
	var (
		request  model.CreateArticleReq
		response pb.CreateArticleResponse
	)
	if err := srv.Transform(req.GetArticle(), &request.BaseArticle); err != nil {
		return nil, err
	}

	result, err := srv.ArticleUseCase.Create(ctx, &request)
	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Article successfully created")
	response.Article = &pb.Article{}

	if err := srv.Transform(result.Article, response.Article); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *ArticleService) UpdateArticle(
	ctx context.Context, req *pb.UpdateArticleRequest,
) (*pb.UpdateArticleResponse, error) {
	var (
		request  model.UpdateArticleReq
		response pb.UpdateArticleResponse
	)
	err := srv.Transform(req.GetArticle(), &request.BaseArticle)
	if err != nil {
		return nil, err
	}

	request.Id = req.Article.Id
	result, err := srv.ArticleUseCase.Update(ctx, &request)

	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Article succesfully updated")
	response.Article = &pb.Article{}
	if err := srv.Transform(result.Article, &response.Article); err != nil {
		return nil, err
	}

	return &response, nil

}

func (srv *ArticleService) GetListArticle(
	ctx context.Context, req *pb.ListArticlesRequest,
) (*pb.ListArticlesResponse, error) {
	var (
		list     pb.ListArticlesResponse
		request  model.GetListArticleReq
		errParse error
	)
	request.Page, request.Order, request.Filter, request.Keyword, ctx, errParse = srv.ParseListParams(ctx, req.GetPagination().GetOffset(), req.GetPagination().GetLimit(), req.GetQuery().GetOrder(), req.GetQuery().GetFilter(), req.GetQuery().GetKeyword())
	if errParse != nil {
		return nil, status.Error(codes.PermissionDenied, errParse.Error())
	}
	findReq := &model.GetListArticleReq{
		Page:    request.Page,
		Filter:  request.Filter,
		Order:   request.Order,
		Keyword: request.Keyword,
	}

	result, err := srv.ArticleUseCase.GetList(ctx, findReq)
	if err != nil {
		return nil, err
	}

	if err := srv.TransformSlice(reflect.ValueOf(result.Data), reflect.ValueOf(&list.Articles).Elem()); err != nil {
		return nil, err
	}

	list.Meta = srv.ResponseOKPagination("Article data retrieved")
	if err := srv.Transform(result.Pagination, list.Meta.Pagination); err != nil {
		return nil, srv.ResponseError(err)
	}

	return &list, nil
}

func (srv *ArticleService) GetArticleById(
	ctx context.Context, req *pb.GetArticleRequest,
) (*pb.GetArticleResponse, error) {
	var (
		request  model.GetIdArticleReq
		response pb.GetArticleResponse
	)

	request.Id = req.GetId()
	result, err := srv.ArticleUseCase.GetById(ctx, &request)
	if err != nil {
		return nil, err
	}

	response.Meta = srv.ResponseOK("Article data retrieved")
	response.Article = &pb.Article{}
	if err := srv.Transform(result.Article, response.Article); err != nil {
		return nil, err
	}

	return &response, nil
}

func (srv *ArticleService) DeleteArticle(
	ctx context.Context, req *pb.DeleteArticleRequest,
) (*pb.DeleteArticleResponse, error) {
	var (
		response pb.DeleteArticleResponse
	)

	result, err := srv.ArticleUseCase.Delete(ctx, &model.DeleteArticleReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	response.Meta = srv.ResponseOK("Article data succesfully deleted")
	response.Article = &pb.Article{}
	if err := srv.Transform(result.Article, response.Article); err != nil {
		return nil, err
	}

	return &response, nil

}

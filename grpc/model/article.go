package model

import (
	"encoding/json"
	"time"

	"cms/grpc/domain"
)

type BaseArticle struct {
	Id        int64           `json:"id"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	AuthorId  string          `json:"author_id"`
	CreatedAt time.Time       `gorm:"<-:create;" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	Tag       json.RawMessage `json:"tag,omitempty"`
}

func (req BaseArticle) ToDomain() *domain.Article {
	return &domain.Article{
		Id:       req.Id,
		Title:    req.Title,
		Content:  req.Content,
		AuthorId: req.AuthorId,
		Tag:      req.Tag,
	}
}

type CreateArticleReq struct {
	*BaseArticle
}

type CreateArticleRes struct {
	*domain.Article
}

type UpdateArticleReq struct {
	Id int64 `json:"id"`
	*BaseArticle
}

type UpdateArticleRes struct {
	*domain.Article
}

type GetListArticleReq struct {
	Page    PaginationParam
	Filter  FilterParams
	Order   OrderParam
	Keyword KeywordParam
}

type GetListArticleRes struct {
	Data       []*domain.Article
	Pagination *Pagination
}

type GetIdArticleReq struct {
	Id int64 `json:"id"`
}

type GetIdArticleRes struct {
	*domain.Article
}

type DeleteArticleReq struct {
	Id int64 `json:"id"`
}

type DeleteArticleRes struct {
	*domain.Article
}

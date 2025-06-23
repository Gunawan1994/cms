package model

import (
	"cms/grpc/domain"
	"time"
)

type BaseCreateUserReq struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (req *BaseCreateUserReq) ToDomain() *domain.User {
	return &domain.User{
		Id:       req.Id,
		Email:    req.Email,
		Password: req.Password,
		Username: req.Username,
	}
}

type CreateUserReq struct {
	*BaseCreateUserReq
}

type CreateUserRes struct {
	*domain.User
}

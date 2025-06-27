package model

import "cms/grpc/domain"

type VerifyCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyCredentialRes struct {
	domain.User
}

type AccessApiCredential struct {
	ApiSecretId  string
	ApiSecretKey string
}

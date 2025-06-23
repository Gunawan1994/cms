package postgres

import (
	"cms/grpc/domain"
	"cms/grpc/module/base/repository"
)

type UserRepository interface {
	repository.BaseRepository[domain.User]
}

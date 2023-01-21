package repository

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
)

type UserRepository interface {
	Authentication(ctx context.Context, username string) (entity.User, error)
	Create(username string, password string, roles []string)
	DeleteAll()
}

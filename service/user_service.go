package service

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
)

type UserService interface {
	Authentication(ctx context.Context, model model.UserModel) entity.User
}

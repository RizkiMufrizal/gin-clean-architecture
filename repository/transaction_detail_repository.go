package repository

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
)

type TransactionDetailRepository interface {
	FindById(ctx context.Context, id string) (entity.TransactionDetail, error)
}

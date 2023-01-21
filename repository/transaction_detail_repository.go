package repository

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
)

type TransactionDetailRepository interface {
	Insert(ctx context.Context, transactionDetails []entity.TransactionDetail) []entity.TransactionDetail
	FindById(ctx context.Context, id string) (entity.TransactionDetail, error)
}

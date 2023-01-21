package service

import (
	"context"
	"github.com/RizkiMufrizal/gin-clean-architecture/model"
)

type TransactionDetailService interface {
	FindById(ctx context.Context, id string) model.TransactionDetailModel
}

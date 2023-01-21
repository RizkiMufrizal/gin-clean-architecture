package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/repository"
	"github.com/jmoiron/sqlx"
)

func NewTransactionDetailRepositoryImpl(DB *sqlx.DB) repository.TransactionDetailRepository {
	return &transactionDetailRepositoryImpl{DB: DB}
}

type transactionDetailRepositoryImpl struct {
	*sqlx.DB
}

func (transactionDetailRepository *transactionDetailRepositoryImpl) Insert(ctx context.Context, transactionDetails []entity.TransactionDetail) []entity.TransactionDetail {
	//begin transaction
	tx, err := transactionDetailRepository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExecContext(ctx, "INSERT INTO tb_transaction_detail (transaction_detail_id, sub_total_price, price, quantity, transaction_id, product_id)VALUES(:transaction_detail_id, :sub_total_price, :price, :quantity, :transaction_id, :product_id)", transactionDetails)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)

	return transactionDetails
}

func (transactionDetailRepository *transactionDetailRepositoryImpl) FindById(ctx context.Context, id string) (entity.TransactionDetail, error) {
	var transactionDetail entity.TransactionDetail
	err := transactionDetailRepository.DB.GetContext(ctx, &transactionDetail, ""+
		"SELECT transaction_detail.transaction_detail_id, transaction_detail.sub_total_price, transaction_detail.price, transaction_detail.quantity, transaction_detail.transaction_id, transaction_detail.product_id, product.name, product.price, product.quantity "+
		"FROM tb_transaction_detail transaction_detail "+
		"INNER JOIN tb_product product ON transaction_detail.product_id = product.product_id "+
		"where transaction_detail.transaction_detail_id = ?", id)
	if err != nil {
		return entity.TransactionDetail{}, errors.New("transaction Detail Not Found")
	}
	return transactionDetail, nil
}

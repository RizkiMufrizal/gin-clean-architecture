package impl

import (
	"context"
	"errors"
	"github.com/RizkiMufrizal/gin-clean-architecture/entity"
	"github.com/RizkiMufrizal/gin-clean-architecture/exception"
	"github.com/RizkiMufrizal/gin-clean-architecture/repository"
	"github.com/jmoiron/sqlx"
)

func NewTransactionRepositoryImpl(DB *sqlx.DB) repository.TransactionRepository {
	return &transactionRepositoryImpl{DB: DB}
}

type transactionRepositoryImpl struct {
	*sqlx.DB
}

func (transactionRepository *transactionRepositoryImpl) Insert(ctx context.Context, transaction entity.Transaction) entity.Transaction {
	//begin transaction
	tx, err := transactionRepository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExecContext(ctx, "INSERT INTO tb_transaction (transaction_id, total_price) VALUES(:transaction_id, :total_price)", transaction)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)

	return transaction
}

func (transactionRepository *transactionRepositoryImpl) Delete(ctx context.Context, transaction entity.Transaction) {
	//begin transaction
	tx, err := transactionRepository.DB.Beginx()
	exception.PanicLogging(err)

	_, err = tx.NamedExecContext(ctx, "DELETE FROM tb_transaction WHERE transaction_id=:transaction_id", transaction)
	if err != nil {
		err := tx.Rollback()
		exception.PanicLogging(err)
	}
	//commit
	err = tx.Commit()
	exception.PanicLogging(err)
}

func (transactionRepository *transactionRepositoryImpl) FindById(ctx context.Context, id string) (entity.Transaction, error) {
	var transaction entity.Transaction
	err := transactionRepository.DB.GetContext(ctx, &transaction, ""+
		"SELECT transaction_detail.transaction_detail_id, transaction_detail.sub_total_price, transaction_detail.price, transaction_detail.quantity, transaction_detail.transaction_id, transaction_detail.product_id, product.name, product.price, product.quantity, transaction.transaction_id, transaction.total_price "+
		"FROM tb_transaction transaction "+
		"INNER JOIN tb_transaction_detail transaction_detail "+
		"ON transaction.transaction_id = transaction_detail.transaction_id "+
		"INNER JOIN tb_product product "+
		"ON transaction_detail.product_id = product.product_id "+
		"where transaction.transaction_id = ?", id)
	if err != nil {
		return entity.Transaction{}, errors.New("transaction Not Found")
	}
	return transaction, nil
}

func (transactionRepository *transactionRepositoryImpl) FindAll(ctx context.Context) []entity.Transaction {
	var transactions []entity.Transaction
	err := transactionRepository.DB.SelectContext(ctx, &transactions, ""+
		"SELECT transaction_detail.transaction_detail_id, transaction_detail.sub_total_price, transaction_detail.price, transaction_detail.quantity, transaction_detail.transaction_id, transaction_detail.product_id, product.name, product.price, product.quantity, transaction.transaction_id, transaction.total_price "+
		"FROM tb_transaction transaction "+
		"INNER JOIN tb_transaction_detail transaction_detail "+
		"ON transaction.transaction_id = transaction_detail.transaction_id "+
		"INNER JOIN tb_product product "+
		"ON transaction_detail.product_id = product.product_id")
	if err != nil {
		exception.PanicLogging(err)
	}
	return transactions
}

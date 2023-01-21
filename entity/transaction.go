package entity

import "github.com/google/uuid"

type Transaction struct {
	Id                 uuid.UUID           `db:"transaction_id"`
	TotalPrice         int64               `db:"total_price"`
	TransactionDetails []TransactionDetail `db:"transaction_details"`
}

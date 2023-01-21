package entity

import "github.com/google/uuid"

type Product struct {
	Id                 uuid.UUID           `db:"product_id"`
	Name               string              `db:"name"`
	Price              int64               `db:"price"`
	Quantity           int32               `db:"quantity"`
	TransactionDetails []TransactionDetail `db:"transaction_details"`
}

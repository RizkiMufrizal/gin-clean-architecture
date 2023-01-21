package entity

import "github.com/google/uuid"

type TransactionDetail struct {
	Id            uuid.UUID   `db:"transaction_detail_id"`
	SubTotalPrice int64       `db:"sub_total_price"`
	Price         int64       `db:"price"`
	Quantity      int32       `db:"quantity"`
	Product       Product     `db:"product"`
	Transaction   Transaction `db:"transaction"`
	TransactionId uuid.UUID   `db:"transaction_id"`
	ProductId     uuid.UUID   `db:"product_id"`
}

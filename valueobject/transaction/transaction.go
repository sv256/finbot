package transaction

import (
	"finbot/valueobject/transaction/type"
	"github.com/google/uuid"
)

type Transaction struct {
	id     uuid.UUID
	_type  _type.TransactionType
	amount float64
	name   string
}

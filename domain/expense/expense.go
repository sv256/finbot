package expense

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrAmountIsZero = errors.New("amount cannot be 0.0")
)

type Expense struct {
	id     uuid.UUID
	amount float64
	desc   string
}

func NewExpense(amount float64, desc string) (Expense, error) {
	if amount == 0.0 {
		return Expense{}, ErrAmountIsZero
	}
	return Expense{
		id:     uuid.New(),
		amount: amount,
		desc:   desc,
	}, nil
}
func (e Expense) GetID() uuid.UUID {
	return e.id
}

func (e Expense) GetAmount() float64 {
	return e.amount
}
func (e Expense) GetDesc() string {
	return e.desc
}

package expense

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrAmountIsZero   = errors.New("amount cannot be 0.0")
	ErrUserIdIsNil    = errors.New("user id cannot be nil")
	ErrAccountIdIsNil = errors.New("account id cannot be nil")
)

// todo make the expense and income from the entity as transaction
type Expense struct {
	id        uuid.UUID
	userId    uuid.UUID
	accountId uuid.UUID
	amount    float64
	desc      string
}

func NewExpense(userId uuid.UUID, accountId uuid.UUID, amount float64, desc string) (Expense, error) {
	if amount == 0.0 {
		return Expense{}, ErrAmountIsZero
	}
	if userId == uuid.Nil {
		return Expense{}, ErrUserIdIsNil
	}
	if accountId == uuid.Nil {
		return Expense{}, ErrAccountIdIsNil
	}
	// todo add the user & account validation
	return Expense{
		id:        uuid.New(),
		userId:    userId,
		accountId: accountId,
		amount:    amount,
		desc:      desc,
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

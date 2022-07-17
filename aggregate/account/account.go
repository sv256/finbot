package account

import (
	"errors"
	"finbot/domain/account/currency"
	"github.com/google/uuid"
)

type Account struct {
	id        uuid.UUID
	userId    uuid.UUID
	accNumber string
	balance   float64
	currency  currency.Currency
}

var (
	ErrMissingValues                        = errors.New("missing values")
	ErrInvalidInitialBalance                = errors.New("initial balance cannot be less than 0")
	ErrAccountNumberAndCurrencyAlreadyExist = errors.New("account number and with defined currency already exists")
)

func NewAccount(userId uuid.UUID, currency currency.Currency, accNumber string, initialBalance float64) (Account, error) {
	if accNumber == "" {
		return Account{}, ErrMissingValues
	}
	if initialBalance < 0 {
		return Account{}, ErrInvalidInitialBalance
	}
	// todo add user existence validation
	return Account{
		id:        uuid.New(),
		userId:    userId,
		accNumber: accNumber,
		balance:   initialBalance,
		currency:  currency,
	}, nil
}

func (a Account) GetID() uuid.UUID {
	return a.id
}

func (a Account) GetAccountNumber() string {
	return a.accNumber
}

func (a Account) GetCurrency() currency.Currency {
	return a.currency
}

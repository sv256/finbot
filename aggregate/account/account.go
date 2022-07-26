package account

import (
	"errors"
	accountEntity "finbot/entity/account"
	userEntity "finbot/entity/user"
	"finbot/olddomain/account/currency"
	"github.com/google/uuid"
)

type Account struct {
	item      *accountEntity.Account
	user      *userEntity.User
	accNumber string
	balance   float64
}

var (
	ErrMissingValues                        = errors.New("missing values")
	ErrInvalidInitialBalance                = errors.New("initial balance cannot be less than 0")
	ErrAccountNumberAndCurrencyAlreadyExist = errors.New("account number and with defined currency already exists")
)

func NewAccount(user *userEntity.User, currency currency.Currency, accNumber string, initialBalance float64) (Account, error) {
	if accNumber == "" {
		return Account{}, ErrMissingValues
	}
	if initialBalance < 0 {
		return Account{}, ErrInvalidInitialBalance
	}
	// todo add user-a existence validation
	newAccount := &accountEntity.Account{
		Id:        uuid.New(),
		AccNumber: accNumber,
		Balance:   initialBalance,
		Currency:  currency,
	}
	return Account{
		item:      newAccount,
		user:      user,
		accNumber: accNumber,
		balance:   initialBalance,
	}, nil
}

func (a Account) GetID() uuid.UUID {
	return a.item.Id
}

func (a Account) GetAccountNumber() string {
	return a.accNumber
}

func (a Account) GetCurrency() currency.Currency {
	return a.item.Currency
}

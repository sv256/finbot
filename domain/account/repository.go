package account

import (
	"errors"
	"finbot/aggregate/account"
	"github.com/google/uuid"
)

var (
	ErrAccountNotFound       = errors.New("the account was not found")
	ErrAccountIdAlreadyExist = errors.New("the account already exist")
)

type AccountRepository interface {
	GetAll() ([]account.Account, error)
	GetByID(uuid uuid.UUID) (account.Account, error)
	Add(account account.Account) error
	Update(account account.Account) error
	Delete(uuid2 uuid.UUID) error
}

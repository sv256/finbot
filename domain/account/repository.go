package account

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrAccountNotFound       = errors.New("the account was not found")
	ErrAccountIdAlreadyExist = errors.New("the account already exist")
)

type AccountRepository interface {
	GetAll() ([]Account, error)
	GetByID(uuid uuid.UUID) (Account, error)
	Add(account Account) error
	Update(account Account) error
	Delete(uuid2 uuid.UUID) error
}

package expense

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrNotExistingExpense  = errors.New("expense is not exists")
	ErrExpenseAlreadyExist = errors.New("expense already exists")
)

type Repository interface {
	GetAll() ([]Expense, error)
	GetById(uuid.UUID) (Expense, error)
	Add(Expense) error
	Update(Expense) error
	Delete(uuid.UUID) error
}

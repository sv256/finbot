package expense

import (
	"errors"
	"finbot/aggregate/expense"
	"github.com/google/uuid"
)

var (
	ErrNotExistingExpense  = errors.New("expense is not exists")
	ErrExpenseAlreadyExist = errors.New("expense already exists")
)

type Repository interface {
	GetAll() ([]expense.Expense, error)
	GetById(uuid.UUID) (expense.Expense, error)
	Add(expense.Expense) error
	Update(expense.Expense) error
	Delete(uuid.UUID) error
}

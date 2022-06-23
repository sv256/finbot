package expense

import "github.com/google/uuid"

type ExpenseRepository interface {
	GetAll() ([]Expense, error)
	GetById(uuid.UUID) (Expense, error)
	Add(Expense) error
	Update(Expense) error
	Delete(uuid.UUID) error
}

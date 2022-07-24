package memory

import (
	"finbot/aggregate/torefactor/expense"
	"github.com/google/uuid"
	"sync"
)

type MemoryExpenseRepository struct {
	expenses map[uuid.UUID]expense.Expense
	sync.Mutex
}

func New() *MemoryExpenseRepository {
	return &MemoryExpenseRepository{
		expenses: make(map[uuid.UUID]expense.Expense),
	}
}

func (mpr *MemoryExpenseRepository) GetAll() ([]expense.Expense, error) {
	var expenses []expense.Expense
	for _, exp := range mpr.expenses {
		expenses = append(expenses, exp)
	}
	return expenses, nil
}

//func (mpr *MemoryExpenseRepository) GetByID(id uuid.UUID) (outcome.Expense, error) {
//	if exp, ok := mpr.expenses[id]; ok {
//		return exp, nil
//	}
//	return outcome.Expense{}, expenses.ErrProductNotFound
//}
//
//func (mpr *MemoryExpenseRepository) Add(newExp outcome.Expense) error {
//	mpr.Lock()
//	defer mpr.Unlock()
//
//	if _, ok := mpr.expenses[newExp.GetID()]; ok {
//		return outcome.ErrProductAlreadyExist
//	}
//
//	mpr.expenses[newExp.GetID()] = newExp
//
//	return nil
//}
//func (mpr *MemoryExpenseRepository) Update(newExp outcome.Expense) error {
//	mpr.Lock()
//	defer mpr.Unlock()
//
//	if _, ok := mpr.expenses[newExp.GetID()]; !ok {
//		return outcome.ErrProductNotFound
//	}
//
//	mpr.expenses[newExp.GetID()] = newExp
//	return nil
//}
//func (mpr *MemoryExpenseRepository) Delete(id uuid.UUID) error {
//	mpr.Lock()
//	defer mpr.Unlock()
//
//	if _, ok := mpr.expenses[id]; !ok {
//		return outcome.ErrProductNotFound
//	}
//	delete(mpr.expenses, id)
//	return nil
//}

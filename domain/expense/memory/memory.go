package memory

import (
	"finbot/aggregate/expense"
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

//func (mpr *MemoryExpenseRepository) GetByID(id uuid.UUID) (expense.Expense, error) {
//	if exp, ok := mpr.expenses[id]; ok {
//		return exp, nil
//	}
//	return expense.Expense{}, expenses.ErrProductNotFound
//}
//
//func (mpr *MemoryExpenseRepository) Add(newExp expense.Expense) error {
//	mpr.Lock()
//	defer mpr.Unlock()
//
//	if _, ok := mpr.expenses[newExp.GetID()]; ok {
//		return expense.ErrProductAlreadyExist
//	}
//
//	mpr.expenses[newExp.GetID()] = newExp
//
//	return nil
//}
//func (mpr *MemoryExpenseRepository) Update(newExp expense.Expense) error {
//	mpr.Lock()
//	defer mpr.Unlock()
//
//	if _, ok := mpr.expenses[newExp.GetID()]; !ok {
//		return expense.ErrProductNotFound
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
//		return expense.ErrProductNotFound
//	}
//	delete(mpr.expenses, id)
//	return nil
//}

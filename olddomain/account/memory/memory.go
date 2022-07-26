package memory

import (
	account2 "finbot/aggregate/account"
	"finbot/olddomain/account"
	"github.com/google/uuid"
	"sync"
)

type MemoryAccountRepository struct {
	accounts                map[uuid.UUID]account2.Account
	accountsByAccountNumber map[string]account2.Account
	sync.Mutex
}

func New() *MemoryAccountRepository {
	return &MemoryAccountRepository{
		accounts: make(map[uuid.UUID]account2.Account),
	}
}

func (mar *MemoryAccountRepository) GetAll() ([]account2.Account, error) {
	// Collect all Products from map
	var accounts []account2.Account
	for _, acc := range mar.accounts {
		accounts = append(accounts, acc)
	}
	return accounts, nil
}

func (mar *MemoryAccountRepository) GetByID(id uuid.UUID) (account2.Account, error) {
	if acc, ok := mar.accounts[id]; ok {
		return acc, nil
	}
	return account2.Account{}, account.ErrAccountNotFound
}

func (mar *MemoryAccountRepository) Add(newAcc account2.Account) error {
	mar.Lock()
	defer mar.Unlock()

	if _, err := mar.accounts[newAcc.GetID()]; err {
		return account.ErrAccountIdAlreadyExist
	} else {

		if accByNum, err := mar.accountsByAccountNumber[newAcc.GetAccountNumber()]; !err && accByNum.GetCurrency() == newAcc.GetCurrency() {
			return account2.ErrAccountNumberAndCurrencyAlreadyExist
		}
	}

	mar.accounts[newAcc.GetID()] = newAcc

	return nil
}
func (mar *MemoryAccountRepository) Delete(id uuid.UUID) error {
	mar.Lock()
	defer mar.Unlock()

	if _, ok := mar.accounts[id]; !ok {
		return account.ErrAccountNotFound
	}
	delete(mar.accounts, id)
	return nil
}

func (mar *MemoryAccountRepository) Update(updateAcc account2.Account) error {
	mar.Lock()
	defer mar.Unlock()

	if _, ok := mar.accounts[updateAcc.GetID()]; !ok {
		return account.ErrAccountNotFound
	}

	mar.accounts[updateAcc.GetID()] = updateAcc
	return nil
}

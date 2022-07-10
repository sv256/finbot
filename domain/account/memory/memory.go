package memory

import (
	"github.com/google/uuid"
	"sync"
)

type MemoryAccountRepository struct {
	accounts                map[uuid.UUID]account.Account
	accountsByAccountNumber map[string]account.Account
	sync.Mutex
}

func New() *MemoryAccountRepository {
	return &MemoryAccountRepository{
		accounts: make(map[uuid.UUID]account.Account),
	}
}

func (mar *MemoryAccountRepository) GetAll() ([]account.Account, error) {
	// Collect all Products from map
	var accounts []account.Account
	for _, account := range mar.accounts {
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (mar *MemoryAccountRepository) GetByID(id uuid.UUID) (account.Account, error) {
	if account, ok := mar.accounts[id]; ok {
		return account, nil
	}
	return account.Account{}, account.ErrAccountNotFound
}

func (mar *MemoryAccountRepository) Add(newAcc account.Account) error {
	mar.Lock()
	defer mar.Unlock()

	if _, err := mar.accounts[newAcc.GetID()]; err {
		return account.ErrAccountIdAlreadyExist
	} else {

		if accByNum, err := mar.accountsByAccountNumber[newAcc.GetAccountNumber()]; !err && accByNum.GetCurrency() == newAcc.GetCurrency {
			return account.ErrAccountNumberAndCurrencyAlreadyExist
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

func (mar *MemoryAccountRepository) Update(updateAcc account.Account) error {
	mar.Lock()
	defer mar.Unlock()

	if _, ok := mar.accounts[updateAcc.GetID()]; !ok {
		return account.ErrProductNotFound
	}

	mar.accounts[updateAcc.GetID()] = updateAcc
	return nil
}

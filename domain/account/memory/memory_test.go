package memory

import (
	"finbot/domain/account"
	"finbot/domain/account/currency"
	"finbot/domain/user"
	"github.com/google/uuid"
	"testing"
)

func TestMemoryAccountRepository_Add(t *testing.T) {
	repo := New()
	u, _ := user.NewUser("TestUserName1")
	acc, err := account.NewAccount(u.GetId(), currency.PLN, "012345", 0.1)
	if err != nil {
		t.Error(err)
	}

	repo.Add(acc)
	if len(repo.accounts) != 1 {
		t.Errorf("Expected 1 acc, got %d", len(repo.accounts))
	}
}

func TestMemoryAccountRepository_Get(t *testing.T) {
	repo := New()
	u, _ := user.NewUser("TestUserName1")
	existingAccount, err := account.NewAccount(u.GetId(), currency.PLN, "012345", 0.1)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingAccount)
	if len(repo.accounts) != 1 {
		t.Errorf("Expected 1 account, got %d", len(repo.accounts))
	}

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Get product by id",
			id:          existingAccount.GetID(),
			expectedErr: nil,
		}, {
			name:        "Get non-existing product by id",
			id:          uuid.New(),
			expectedErr: account.ErrAccountNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.GetByID(tc.id)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}
}
func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	u, _ := user.NewUser("TestUserName1")
	existingAccount, err := account.NewAccount(u.GetId(), currency.PLN, "012345", 1.1)
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingAccount)
	if len(repo.accounts) != 1 {
		t.Errorf("Expected 1 product, got %d", len(repo.accounts))
	}

	err = repo.Delete(existingAccount.GetID())
	if err != nil {
		t.Error(err)
	}
	if len(repo.accounts) != 0 {
		t.Errorf("Expected 0 products, got %d", len(repo.accounts))
	}
}

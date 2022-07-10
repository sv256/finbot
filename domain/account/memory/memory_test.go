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
	acc, err := account.NewAccount(uuid.New(), u.GetId(), currency.PLN, "012345")
	if err != nil {
		t.Error(err)
	}

	repo.Add(acc)
	if len(repo.accounts) != 1 {
		t.Errorf("Expected 1 acc, got %d", len(repo.account))
	}
}

func TestMemoryAccountRepository_Get(t *testing.T) {
	repo := New()
	u, _ := user.NewUser("TestUserName1")
	existingAccount, err := account.NewAccount(uuid.New(), u.GetId(), currency.PLN, "012345")
	if err != nil {
		t.Error(err)
	}

	repo.Add(existingAccount)
	if len(repo.accounts) != 1 {
		t.Errorf("Expected 1 account, got %d", len(repo.existingAccount))
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
	existingAccount, err := account.NewAccount(uuid.New(), u.GetId(), currency.PLN, "012345")
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

//func TestMemoryAccountRepository_InitialBalance(t *testing.T) {
//	repo := New()
//	u, _ := user.NewUser("TestUserName1")
//	existingAccount, err := aggregate.NewAccount(uuid.New(), u, currency.PLN, "012345")
//
//	if err != nil {
//		t.Error(err)
//	}
//
//	repo.Add(existingAccount)
//	if len(repo.accounts) != 1 {
//		t.Errorf("Expected 1 account, got %d", len(repo.existingAccount))
//	}
//	type testCase struct {
//		name        string
//		id          uuid.UUID
//		balance     float64
//		expectedErr error
//	}
//
//	testCases := []testCase{
//		{
//			name:        "Account init balance",
//			id:          existingProd.GetID(),
//			balance:     1000.0,
//			expectedErr: nil,
//		}, {
//			name:        "Non-existing account to assign a balance",
//			id:          uuid.New(),
//			balance:     1000.0,
//			expectedErr: product.ErrAccountNotFound,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			_, err := repo.GetByID(tc.id)
//			if err != tc.expectedErr {
//				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
//			}
//
//		})
//	}
//
//	_, err = aggregate.InitBalance(account.GetUUID(), 1000.0)
//
//}

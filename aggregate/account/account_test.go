package account_test

import (
	"finbot/aggregate/account"
	"finbot/domain/account/currency"
	"github.com/google/uuid"
	"testing"
)

func TestAccount_NewAccount(t *testing.T) {
	type testCase struct {
		test           string
		userId         uuid.UUID
		accNumber      string
		initialBalance float64
		expectedErr    error
	}

	testCases := []testCase{
		{
			test:        "should return error if account number is empty",
			userId:      uuid.New(),
			expectedErr: account.ErrMissingValues,
		},
		{
			test:           "Wrong initial balance",
			userId:         uuid.New(),
			accNumber:      "012345",
			initialBalance: -1.0,
			expectedErr:    account.ErrInvalidInitialBalance,
		},
		{
			test:           "Valid Test",
			userId:         uuid.New(),
			accNumber:      "012345",
			initialBalance: 1.0,
			expectedErr:    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := account.NewAccount(tc.userId, currency.PLN, tc.accNumber, tc.initialBalance)
			if err != tc.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}

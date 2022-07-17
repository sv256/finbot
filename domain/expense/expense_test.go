package expense_test

import (
	"finbot/domain/expense"
	"github.com/google/uuid"
	"testing"
)

func TestNewExpense(t *testing.T) {
	type testCase struct {
		test        string
		amount      float64
		desc        string
		expectedErr error
		userId      uuid.UUID
		accountId   uuid.UUID
	}

	testCases := []testCase{
		{
			test:        "Should return error if amount is 0.0",
			amount:      0.0,
			expectedErr: expense.ErrAmountIsZero,
		},
		{
			test:        "Should return error for null userId",
			amount:      0.1,
			desc:        "test description",
			userId:      uuid.Nil,
			accountId:   uuid.New(),
			expectedErr: expense.ErrUserIdIsNil,
		},
		{
			test:        "Should return error for null accountId",
			amount:      0.1,
			desc:        "test description",
			userId:      uuid.New(),
			accountId:   uuid.Nil,
			expectedErr: expense.ErrAccountIdIsNil,
		},
		{
			test:        "Should return correct object with description",
			amount:      0.1,
			desc:        "test description",
			userId:      uuid.New(),
			accountId:   uuid.New(),
			expectedErr: nil,
		},
		{
			test:        "Should return correct object without",
			amount:      0.1,
			userId:      uuid.New(),
			accountId:   uuid.New(),
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			obj, err := expense.NewExpense(tc.userId, tc.accountId, tc.amount, tc.desc)
			if tc.expectedErr != nil {
				if err != tc.expectedErr {
					t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
				}
			} else {
				if tc.desc != "" && tc.desc != obj.GetDesc() {
					t.Errorf("Expected descr %v, got %v", tc.desc, obj.GetDesc())
				}
				if tc.amount != obj.GetAmount() {
					t.Errorf("Expected amount %v, got %v", tc.amount, obj.GetAmount())
				}
			}
		})
	}

}

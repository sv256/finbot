package expense_test

import (
	"finbot/domain/expense"
	"testing"
)

func TestNewExpense(t *testing.T) {
	type testCase struct {
		test        string
		amount      float64
		desc        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Should return error if amount is 0.0",
			amount:      0.0,
			expectedErr: expense.ErrAmountIsZero,
		},
		{
			test:        "Should return correct object with description",
			amount:      0.1,
			desc:        "test description",
			expectedErr: nil,
		},
		{
			test:        "Should return correct object without",
			amount:      0.1,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			obj, err := expense.NewExpense(tc.amount, tc.desc)
			if tc.desc != "" && tc.desc != obj.GetDesc() {
				t.Errorf("Expected descr %v, got %v", tc.desc, obj.GetDesc())
			}
			if tc.amount != obj.GetAmount() {
				t.Errorf("Expected amount %v, got %v", tc.amount, obj.GetAmount())
			}
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}

}

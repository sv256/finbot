package user_test

import (
	"finbot/aggregate/user"
	"github.com/google/uuid"
	"testing"
)

func TestNewUser(t *testing.T) {
	type testCase struct {
		test        string
		id          uuid.UUID
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Should create a user-a with defined name",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			name:        "Andrew",
			expectedErr: nil,
		},
		{
			test:        "Should throw an error for a user-a without defined name",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			name:        "",
			expectedErr: user.ErrUserHasNoName,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			obj, err := user.NewUser(tc.name)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if !obj.IsEmpty() {
				if tc.name != obj.GetName() {
					t.Errorf("Expected name %v, got %v", tc.name, obj.GetName())
				}
			}

		})
	}

}

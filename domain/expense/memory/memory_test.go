package memory_test

//func TestMemoryExpenseRepository_Add(t *testing.T) {
//	repository := New()
//
//	object, err := expense.NewExpense(1.1, "shirt")
//
//	if err != nil {
//		t.Errorf(err)
//	}
//	repository.Add(object)
//
//	if len(repository.expenses) != 1 {
//		t.Errorf("Expected size of expenses : 1, got %v", len(repository.expenses))
//	}
//}
//
//func TestMemoryExpenseRepository_Get(t *testing.T) {
//	repository := New()
//	existingObject, err := expense.NewExpense(1.1, "shirt")
//	if err != nil {
//		t.Errorf(err)
//	}
//	repository.Add(existingObject)
//
//	if len(repository.expenses) != 1 {
//		t.Errorf("Expected size of expenses : 1, got %v", len(repository.expenses))
//	}
//
//	type testCase struct {
//		name        string
//		amount      float64
//		desc        string
//		id          uuid.UUID
//		expectedErr error
//	}
//
//	testCases := []testCase{
//		{
//			name:        "Gets the expense by id",
//			id:          existingObject.GetID(),
//			amount:      existingObject.GetAmount(),
//			desc:        existingObject.GetDesc(),
//			expectedErr: nil,
//		},
//		{
//			name:        "Gets non-existing expense by id",
//			id:          uuid.New(),
//			amount:      1.1,
//			desc:        existingObject.GetDesc(),
//			expectedErr: existingObject.ErrNotExistingExpense,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			_, err := repository.GetByID(tc.id)
//			if err != tc.expectedErr {
//				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
//			}
//
//		})
//	}
//
//}
//
//func TestMemoryExpenseRepository_Delete(t *testing.T) {
//	repository := New()
//	existingObject, err := expense.NewExpense(1.1, "shirt")
//	if err != nil {
//		t.Errorf(err)
//	}
//	repository.Add(existingObject)
//
//	if len(repository.expenses) != 1 {
//		t.Errorf("Expected size of expenses : 1, got %d", len(repository.expenses))
//	}
//
//	err = repository.Delete(existingObject.GetID())
//	if err != nil {
//		t.Error(err)
//	}
//
//	if len(repository.expenses) != 0 {
//		t.Errorf("Expected 0 expenses, got %d", len(repository.expenses))
//	}
//
//}

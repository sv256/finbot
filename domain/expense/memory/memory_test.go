package memory_test

//func TestMemory_GetUser(t *testing.T) {
//	type testCase struct {
//		id          uuid.UUID
//		name        string
//		expectedErr error
//	}
//
//	u, err := user.NewUser("Arnie")
//	err != nil{
//		t.Fatal(err)
//	}
//	id := u.GetId()
//	repo := expense.Repository{
//		users: map[uuid.UUID]user.User{
//			id: u
//		},
//	}
//
//	testCases := []testCase{
//		{
//			name: "No user found by id",
//			id: 	uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
//			expectedErr: user.ErrUserNotFound,
//		},{
//			name: "User found by id",
//			id: 	id,
//			expectedErr: nil,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//
//			_, err := repo.Get(tc.id)
//			if err != tc.expectedErr {
//				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
//			}
//		})
//	}
//
//}

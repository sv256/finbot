package user

import (
	"finbot/aggregate/user"
	"strconv"
	"testing"
)

func TestSingleUser_NewUserService(t *testing.T) {
	us, err := NewUserService(
		WithMemoryUserRepository(),
	)

	if err != nil {
		t.Error(err)
	}

	u, err := user.NewUser("test1")

	_, err = us.Add(u)

	if err != nil {
		t.Error(err)
	}

}

func TestMultiUser_NewUserService(t *testing.T) {
	us, err := NewUserService(
		WithMemoryUserRepository(),
	)

	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		u, err := user.NewUser("test" + strconv.Itoa(i))
		_, err = us.Add(u)

		if err != nil {
			t.Error(err)
		}
	}
}

func TestMultiUser_NewUserService_expectError(t *testing.T) {
	us, err := NewUserService(
		WithMemoryUserRepository(),
	)

	if err != nil {
		t.Error(err)
	}
	u, err := us.NewUser("test")
	_, err = us.Add(u)
	if err != nil {
		t.Error(err)
	}
	_, err = us.Add(u)
	if err == nil {
		t.Fail()
	}
}

package user

import (
	"errors"
	accountEntity "finbot/entity/account"
	userEntity "finbot/entity/user"
	valueobject "finbot/valueobject/transaction"
	"github.com/google/uuid"
)

var (
	ErrUserHasNoName = errors.New("user should contain a name")
)

type User struct {
	user         *userEntity.User
	accounts     []*accountEntity.Account
	transactions []valueobject.Transaction
}

func NewUser(name string) (User, error) {
	if name == "" {
		return User{}, ErrUserHasNoName
	}

	// todo validation for existing user

	newUser := &userEntity.User{
		Id:   uuid.New(),
		Name: name,
	}

	return User{
		user:         newUser,
		accounts:     make([]*accountEntity.Account, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (u User) GetName() string {
	return u.user.Name
}

func (u User) GetId() uuid.UUID {
	return u.user.Id
}

func (u User) IsEmpty() bool {
	return u.user == nil
}

package user

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrUserHasNoName = errors.New("user should contain a name")
)

type User struct {
	id   uuid.UUID
	name string
}

func NewUser(name string) (User, error) {
	if name == "" {
		return User{}, ErrUserHasNoName
	}
	return User{
		id:   uuid.New(),
		name: name,
	}, nil
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetId() uuid.UUID {
	return u.id
}

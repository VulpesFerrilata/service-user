package model

import (
	"github.com/VulpesFerrilata/user/internal/domain/datamodel"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(username string, password string) (*User, error) {
	user := new(User)
	user.username = username

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "model.NewUser")
	}
	user.hashPassword = hashPassword
	return user, nil
}

func EmptyUser() *User {
	return new(User)
}

type User struct {
	id           uint
	username     string
	hashPassword []byte
}

func (u User) GetId() uint {
	return u.id
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetHashPassword() []byte {
	return u.hashPassword
}

func (u *User) Persist(f func(user *datamodel.User) error) error {
	user := new(datamodel.User)
	user.ID = u.id
	user.Username = u.username
	user.HashPassword = u.hashPassword
	if err := f(user); err != nil {
		return errors.Wrap(err, "model.User.Persist")
	}
	u.id = user.ID
	u.username = user.Username
	u.hashPassword = user.HashPassword
	return nil
}

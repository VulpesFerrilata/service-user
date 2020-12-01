package model

import (
	"github.com/VulpesFerrilata/user/internal/domain/datamodel"
	"golang.org/x/crypto/bcrypt"
)

func FromUser(u *datamodel.User) *User {
	user := new(User)
	user.User = *u
	return user
}

type User struct {
	datamodel.User
	Password string
}

func (u *User) EncryptPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.HashPassword = hashPassword
	return nil
}

func (u *User) ValidatePassword() error {
	return bcrypt.CompareHashAndPassword(u.HashPassword, []byte(u.Password))
}

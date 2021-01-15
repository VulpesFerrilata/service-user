package datamodel

import (
	"github.com/VulpesFerrilata/user/internal/domain/model"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(username string, password string) (*User, error) {
	user := new(User)
	user.username = username

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "model.NewUser")
	}
	user.id = id

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.Wrap(err, "model.NewUser")
	}
	user.hashPassword = hashPassword

	return user, nil
}

func NewUserFromUserModel(userModel *model.User) *User {
	user := new(User)
	user.id = userModel.ID
	user.username = userModel.Username
	user.hashPassword = userModel.HashPassword
	user.displayName = userModel.DisplayName
	user.email = userModel.Email
	return user
}

type User struct {
	id           uuid.UUID
	username     string
	hashPassword []byte
	displayName  string
	email        string
}

func (u User) GetId() uuid.UUID {
	return u.id
}

func (u User) GetUsername() string {
	return u.username
}

func (u User) GetHashPassword() []byte {
	return u.hashPassword
}

func (u User) GetDisplayName() string {
	return u.displayName
}

func (u User) GetEmail() string {
	return u.email
}

func (u User) ToModel() *model.User {
	userModel := new(model.User)
	userModel.ID = u.id
	userModel.Username = u.username
	userModel.HashPassword = u.hashPassword
	userModel.DisplayName = u.displayName
	userModel.Email = u.email
	return userModel
}

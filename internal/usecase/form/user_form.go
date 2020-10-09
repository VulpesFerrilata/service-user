package form

import (
	"github.com/VulpesFerrilata/user/internal/domain/model"
)

type UserForm struct {
	ID int `name:"id" validate:"required,gt=0"`
}

func (uf UserForm) ToUser() (*model.User, error) {
	user := new(model.User)
	user.ID = uint(uf.ID)
	return user, nil
}

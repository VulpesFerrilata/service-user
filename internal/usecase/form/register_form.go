package form

import "github.com/VulpesFerrilata/user/internal/domain/model"

type RegisterForm struct {
	*LoginForm
	RepeatPassword string `name:"repeat password" validate:"eqfield=Password"`
}

func (rf RegisterForm) ToUser() (*model.User, error) {
	user, err := rf.LoginForm.ToUser()
	if err != nil {
		return nil, err
	}
	return user, nil
}

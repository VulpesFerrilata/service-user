package request

import "github.com/VulpesFerrilata/user/internal/usecase/form"

type RegisterRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

func (rr RegisterRequest) ToInteractorRegisterForm() *form.RegisterForm {
	registerForm := new(form.RegisterForm)
	registerForm.LoginForm = new(form.LoginForm)
	registerForm.Username = rr.Username
	registerForm.Password = rr.Password
	registerForm.RepeatPassword = rr.RepeatPassword
	return registerForm
}

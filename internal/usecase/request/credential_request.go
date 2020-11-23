package request

import "github.com/VulpesFerrilata/user/internal/domain/model"

type CredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (cr CredentialRequest) ToUser() *model.User {
	user := new(model.User)
	user.Username = cr.Username
	user.Password = cr.Password
	return user
}

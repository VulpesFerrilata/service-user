package request

import "github.com/VulpesFerrilata/user/internal/domain/datamodel"

type CredentialRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (cr CredentialRequest) ToUser() *datamodel.User {
	user := new(datamodel.User)
	user.Username = cr.Username
	return user
}

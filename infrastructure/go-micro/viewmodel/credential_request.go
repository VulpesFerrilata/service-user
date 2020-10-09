package viewmodel

import (
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/user/internal/usecase/form"
)

func NewCredentialRequest(credentialRequestPb *user.CredentialRequest) *CredentialRequest {
	return &CredentialRequest{
		credentialRequestPb: credentialRequestPb,
	}
}

type CredentialRequest struct {
	credentialRequestPb *user.CredentialRequest
}

func (cr CredentialRequest) ToLoginForm() *form.LoginForm {
	loginForm := new(form.LoginForm)
	loginForm.Username = cr.credentialRequestPb.GetUsername()
	loginForm.Password = cr.credentialRequestPb.GetPassword()
	return loginForm
}

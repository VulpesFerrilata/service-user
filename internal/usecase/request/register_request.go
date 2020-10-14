package request

type RegisterRequest struct {
	*CredentialRequest
	RepeatPassword string `json:"repeatPassword" validate:"required,eqfield=Password"`
}

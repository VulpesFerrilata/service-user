package business_rule_error

import (
	"github.com/VulpesFerrilata/library/pkg/app_error"
	ut "github.com/go-playground/universal-translator"
)

func NewIncorrectPasswordError() app_error.BusinessRuleError {
	return &IncorrectPasswordError{}
}

type IncorrectPasswordError struct{}

func (ipe IncorrectPasswordError) Error() string {
	return "password is incorrect"
}

func (ipe IncorrectPasswordError) Translate(trans ut.Translator) (string, error) {
	return trans.T("incorrect-password-error")
}

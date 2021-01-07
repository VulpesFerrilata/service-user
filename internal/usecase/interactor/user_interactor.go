package interactor

import (
	"context"

	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/user/internal/domain/service"
	"github.com/VulpesFerrilata/user/internal/usecase/request"
	"github.com/VulpesFerrilata/user/internal/usecase/response"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type UserInteractor interface {
	GetUserById(ctx context.Context, userRequest *request.UserRequest) (*response.UserResponse, error)
	GetUserByCredential(ctx context.Context, credentialRequest *request.CredentialRequest) (*response.UserResponse, error)
	Register(ctx context.Context, registerForm *request.RegisterRequest) (*response.UserResponse, error)
}

func NewUserInteractor(validate *validator.Validate,
	userService service.UserService) UserInteractor {
	return &userInteractor{
		validate:    validate,
		userService: userService,
	}
}

type userInteractor struct {
	validate    *validator.Validate
	userService service.UserService
}

func (ui userInteractor) GetUserById(ctx context.Context, userRequest *request.UserRequest) (*response.UserResponse, error) {
	if err := ui.validate.StructCtx(ctx, userRequest); err != nil {
		if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			err = app_error.NewValidationError(fieldErrors)
		}
		return nil, errors.Wrap(err, "interactor.UserInteractor.GetUserById")
	}

	user, err := ui.userService.GetUserRepository().GetById(ctx, userRequest.ID)
	if err != nil {
		return nil, errors.Wrap(err, "interactor.UserInteractor.GetUserById")
	}

	return response.NewUserResponse(user), nil
}

func (ui userInteractor) GetUserByCredential(ctx context.Context, credentialRequest *request.CredentialRequest) (*response.UserResponse, error) {
	if err := ui.validate.StructCtx(ctx, credentialRequest); err != nil {
		if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			err = app_error.NewValidationError(fieldErrors)
		}
		return nil, errors.Wrap(err, "interactor.UserInteractor.GetUserByCredential")
	}

	if err := ui.userService.ValidateCredential(ctx, credentialRequest.Username, credentialRequest.Password); err != nil {
		return nil, errors.Wrap(err, "interactor.UserInteractor.GetUserByCredential")
	}

	user, err := ui.userService.GetUserRepository().GetByUsername(ctx, credentialRequest.Username)
	if err != nil {
		return nil, errors.Wrap(err, "interactor.UserInteractor.GetUserByCredential")
	}

	return response.NewUserResponse(user), nil
}

func (ui userInteractor) Register(ctx context.Context, registerRequest *request.RegisterRequest) (*response.UserResponse, error) {
	if err := ui.validate.StructCtx(ctx, registerRequest); err != nil {
		if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
			err = app_error.NewValidationError(fieldErrors)
		}
		return nil, errors.Wrap(err, "interactor.UserInteractor.Register")
	}

	user, err := ui.userService.NewUser(ctx, registerRequest.Username, registerRequest.Password)
	if err != nil {
		return nil, errors.Wrap(err, "interactor.UserInteractor.Register")
	}

	if err := ui.userService.GetUserRepository().Insert(ctx, user); err != nil {
		return nil, errors.Wrap(err, "interactor.UserInteractor.Register")
	}

	return response.NewUserResponse(user), nil
}

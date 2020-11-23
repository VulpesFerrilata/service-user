package interactor

import (
	"context"

	"github.com/VulpesFerrilata/library/pkg/validator"
	"github.com/VulpesFerrilata/user/internal/domain/service"
	"github.com/VulpesFerrilata/user/internal/usecase/request"
	"github.com/VulpesFerrilata/user/internal/usecase/response"
)

type UserInteractor interface {
	GetUserById(ctx context.Context, userRequest *request.UserRequest) (*response.UserResponse, error)
	GetUserByCredential(ctx context.Context, credentialRequest *request.CredentialRequest) (*response.UserResponse, error)
	Register(ctx context.Context, registerForm *request.RegisterRequest) (*response.UserResponse, error)
}

func NewUserInteractor(validate validator.Validate,
	userService service.UserService) UserInteractor {
	return &userInteractor{
		validate:    validate,
		userService: userService,
	}
}

type userInteractor struct {
	validate    validator.Validate
	userService service.UserService
}

func (ui userInteractor) GetUserById(ctx context.Context, userRequest *request.UserRequest) (*response.UserResponse, error) {
	if err := ui.validate.Struct(ctx, userRequest); err != nil {
		return nil, err
	}

	user, err := ui.userService.GetUserRepository().GetById(ctx, userRequest.ID)
	if err != nil {
		return nil, err
	}

	return response.NewUserResponse(user), nil
}

func (ui userInteractor) GetUserByCredential(ctx context.Context, credentialRequest *request.CredentialRequest) (*response.UserResponse, error) {
	if err := ui.validate.Struct(ctx, credentialRequest); err != nil {
		return nil, err
	}

	user := credentialRequest.ToUser()
	if err := ui.userService.ValidateCredential(ctx, user); err != nil {
		return nil, err
	}

	user, err := ui.userService.GetUserRepository().GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	return response.NewUserResponse(user), nil
}

func (ui userInteractor) Register(ctx context.Context, registerRequest *request.RegisterRequest) (*response.UserResponse, error) {
	if err := ui.validate.Struct(ctx, registerRequest); err != nil {
		return nil, err
	}

	user := registerRequest.ToUser()
	if err := ui.userService.Create(ctx, user); err != nil {
		return nil, err
	}

	return response.NewUserResponse(user), nil
}

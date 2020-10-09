package interactor

import (
	"context"

	"github.com/VulpesFerrilata/user/internal/domain/service"
	"github.com/VulpesFerrilata/user/internal/usecase/adapter"
	"github.com/VulpesFerrilata/user/internal/usecase/dto"
	"github.com/VulpesFerrilata/user/internal/usecase/form"
)

type UserInteractor interface {
	GetUserById(ctx context.Context, userForm *form.UserForm) (*dto.UserDTO, error)
	GetUserByCredential(ctx context.Context, loginForm *form.LoginForm) (*dto.UserDTO, error)
	Register(ctx context.Context, registerForm *form.RegisterForm) (*dto.UserDTO, error)
}

func NewUserInteractor(userService service.UserService,
	userAdapter adapter.UserAdapter) UserInteractor {
	return &userInteractor{
		userService: userService,
		userAdapter: userAdapter,
	}
}

type userInteractor struct {
	userService service.UserService
	userAdapter adapter.UserAdapter
}

func (ui userInteractor) GetUserById(ctx context.Context, userForm *form.UserForm) (*dto.UserDTO, error) {
	user, err := ui.userAdapter.ParseUser(ctx, userForm)
	if err != nil {
		return nil, err
	}

	user, err = ui.userService.GetUserRepository().GetById(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return ui.userAdapter.ResponseUser(ctx, user)
}

func (ui userInteractor) GetUserByCredential(ctx context.Context, loginForm *form.LoginForm) (*dto.UserDTO, error) {
	user, err := ui.userAdapter.ParseLogin(ctx, loginForm)
	if err != nil {
		return nil, err
	}

	if err := ui.userService.ValidateLogin(ctx, user, loginForm.Password); err != nil {
		return nil, err
	}

	user, err = ui.userService.GetUserRepository().GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	return ui.userAdapter.ResponseUser(ctx, user)
}

func (ui userInteractor) Register(ctx context.Context, registerForm *form.RegisterForm) (*dto.UserDTO, error) {
	user, err := ui.userAdapter.ParseRegister(ctx, registerForm)
	if err != nil {
		return nil, err
	}

	if err := ui.userService.Create(ctx, user); err != nil {
		return nil, err
	}

	return ui.userAdapter.ResponseUser(ctx, user)
}

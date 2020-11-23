package service

import (
	"context"

	"github.com/VulpesFerrilata/user/internal/domain/model"
	"github.com/VulpesFerrilata/user/internal/domain/repository"

	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type UserService interface {
	GetUserRepository() repository.SafeUserRepository
	ValidateCredential(ctx context.Context, user *model.User) error
	Create(ctx context.Context, user *model.User) error
}

func NewUserService(userRepository repository.UserRepository,
	translatorMiddleware *middleware.TranslatorMiddleware) UserService {
	return &userService{
		userRepository:       userRepository,
		translatorMiddleware: translatorMiddleware,
	}
}

type userService struct {
	userRepository       repository.UserRepository
	translatorMiddleware *middleware.TranslatorMiddleware
}

func (us userService) GetUserRepository() repository.SafeUserRepository {
	return us.userRepository
}

func (us userService) ValidateCredential(ctx context.Context, user *model.User) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()

	userDB, err := us.userRepository.GetByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	user.User = userDB.User

	if err := user.ValidatePassword(); err != nil {
		fieldErr, _ := trans.T("validation-invalid", "password")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}

	return nil
}

func (us userService) validate(ctx context.Context, user *model.User) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()

	count, err := us.userRepository.CountByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if count > 0 {
		fieldErr, _ := trans.T("validation-already-exists", "username")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}
	return nil
}

func (us userService) Create(ctx context.Context, user *model.User) error {
	if err := user.EncryptPassword(); err != nil {
		return err
	}

	if err := us.validate(ctx, user); err != nil {
		return err
	}

	return us.userRepository.Insert(ctx, user)
}

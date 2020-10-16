package service

import (
	"context"

	"github.com/VulpesFerrilata/user/internal/domain/datamodel"
	"github.com/VulpesFerrilata/user/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"

	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"github.com/VulpesFerrilata/library/pkg/middleware"
)

type UserService interface {
	GetUserRepository() repository.SafeUserRepository
	ValidateCredential(ctx context.Context, username string, plainPassword string) error
	Create(ctx context.Context, user *datamodel.User, plainPassword string) error
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

func (us userService) ValidateCredential(ctx context.Context, username string, plainPassword string) error {
	trans := us.translatorMiddleware.Get(ctx)
	validationErrs := server_errors.NewValidationError()

	user, err := us.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(user.HashPassword, []byte(plainPassword)); err != nil {
		fieldErr, _ := trans.T("validation-invalid", "password")
		validationErrs.WithFieldError(fieldErr)
	}

	if validationErrs.HasErrors() {
		return validationErrs
	}

	return nil
}

func (us userService) validate(ctx context.Context, user *datamodel.User) error {
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

func (us userService) Create(ctx context.Context, user *datamodel.User, plainPassword string) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashPassword = hashPassword

	if err := us.validate(ctx, user); err != nil {
		return err
	}
	return us.userRepository.Insert(ctx, user)
}

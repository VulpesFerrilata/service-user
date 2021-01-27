package service

import (
	"context"

	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/user/internal/domain/datamodel"
	"github.com/VulpesFerrilata/user/internal/domain/repository"
	"github.com/VulpesFerrilata/user/internal/pkg/app_error/business_rule_error"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUserRepository() repository.UserRepository
	NewUser(ctx context.Context, username string, password string) (*datamodel.User, error)
	ValidateCredential(ctx context.Context, username string, password string) error
}

func NewUserService(
	userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

type userService struct {
	userRepository repository.UserRepository
}

func (us userService) GetUserRepository() repository.UserRepository {
	return us.userRepository
}

func (us userService) isExists(ctx context.Context, username string) (bool, error) {
	_, err := us.userRepository.GetByUsername(ctx, username)
	if err != nil {
		if _, ok := errors.Cause(err).(*app_error.NotFoundError); ok {
			return false, nil
		}
		return false, errors.Wrap(err, "service.UserService.isExists")
	}
	return true, nil
}

func (us userService) NewUser(ctx context.Context, username string, password string) (*datamodel.User, error) {
	isExists, err := us.isExists(ctx, username)
	if err != nil {
		return nil, errors.Wrap(err, "service.UserService.NewUser")
	}
	if isExists {
		return nil, app_error.NewAlreadyExistsError("username")
	}
	user, err := datamodel.NewUser(username, password)
	return user, errors.Wrap(err, "service.NewUser")
}

func (us *userService) ValidateCredential(ctx context.Context, username string, password string) error {
	user, err := us.userRepository.GetByUsername(ctx, username)
	if err != nil {
		return errors.Wrap(err, "service.UserService.ValidateCredential")
	}

	if err := bcrypt.CompareHashAndPassword(user.GetHashPassword(), []byte(password)); err != nil {
		var businessRuleErrors app_error.BusinessRuleErrors
		incorrectPasswordError := business_rule_error.NewIncorrectPasswordError()
		businessRuleErrors = append(businessRuleErrors, incorrectPasswordError)
		return businessRuleErrors
	}

	return nil
}

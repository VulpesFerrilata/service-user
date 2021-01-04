package repository

import (
	"context"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"

	"gorm.io/gorm"

	"github.com/VulpesFerrilata/library/pkg/app_error"
	"github.com/VulpesFerrilata/library/pkg/middleware"
	"github.com/VulpesFerrilata/user/internal/domain/datamodel"
	"github.com/VulpesFerrilata/user/internal/domain/model"
)

type SafeUserRepository interface {
	GetById(ctx context.Context, id int) (*datamodel.User, error)
	GetByUsername(ctx context.Context, username string) (*datamodel.User, error)
}

type UserRepository interface {
	SafeUserRepository
	Insert(context.Context, *datamodel.User) error
}

func NewUserRepository(transactionMiddleware *middleware.TransactionMiddleware,
	validate *validator.Validate) UserRepository {
	return &userRepository{
		transactionMiddleware: transactionMiddleware,
		validate:              validate,
	}
}

type userRepository struct {
	transactionMiddleware *middleware.TransactionMiddleware
	validate              *validator.Validate
}

func (ur userRepository) GetById(ctx context.Context, id int) (*datamodel.User, error) {
	userModel := new(model.User)
	err := ur.transactionMiddleware.Get(ctx).First(userModel, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = app_error.NewNotFoundError("user")
	}
	return datamodel.NewUserFromUserModel(userModel), errors.Wrap(err, "repository.UserRepository.GetById")
}

func (ur userRepository) GetByUsername(ctx context.Context, username string) (*datamodel.User, error) {
	userModel := new(model.User)
	err := ur.transactionMiddleware.Get(ctx).First(userModel, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = app_error.NewNotFoundError("user")
	}
	return datamodel.NewUserFromUserModel(userModel), errors.Wrap(err, "repository.UserRepository.GetByUsername")
}

func (ur userRepository) Insert(ctx context.Context, user *datamodel.User) error {
	return user.Persist(func(userModel *model.User) error {
		if err := ur.validate.StructCtx(ctx, userModel); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewEntityValidationError(userModel, fieldErrors)
			}
			return errors.Wrap(err, "repository.UserRepository.Insert")
		}

		err := ur.transactionMiddleware.Get(ctx).Create(userModel).Error
		return errors.Wrap(err, "repository.UserRepository.Insert")
	})

}

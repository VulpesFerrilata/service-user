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
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
}

type UserRepository interface {
	SafeUserRepository
	Insert(context.Context, *model.User) error
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

func (ur userRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	user := model.EmptyUser()
	return user, user.Persist(func(user *datamodel.User) error {
		err := ur.transactionMiddleware.Get(ctx).First(user, id).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = app_error.NewNotFoundError("user")
		}
		return errors.Wrap(err, "repository.UserRepository.GetById")
	})
}

func (ur userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := model.EmptyUser()
	return user, user.Persist(func(user *datamodel.User) error {
		err := ur.transactionMiddleware.Get(ctx).First(user, "username = ?", username).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = app_error.NewNotFoundError("user")
		}
		return errors.Wrap(err, "repository.UserRepository.GetByUsername")
	})
}

func (ur userRepository) Insert(ctx context.Context, user *model.User) error {
	return user.Persist(func(user *datamodel.User) error {
		if err := ur.validate.StructCtx(ctx, user); err != nil {
			if fieldErrors, ok := errors.Cause(err).(validator.ValidationErrors); ok {
				err = app_error.NewValidationError(app_error.EntityValidation, "user", fieldErrors)
			}
			return errors.Wrap(err, "repository.UserRepository.Insert")
		}

		err := ur.transactionMiddleware.Get(ctx).Create(user).Error
		return errors.Wrap(err, "repository.UserRepository.Insert")
	})
}

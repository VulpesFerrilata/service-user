package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/VulpesFerrilata/library/pkg/db"
	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
	"github.com/VulpesFerrilata/user/internal/domain/datamodel"
)

type SafeUserRepository interface {
	CountByUsername(ctx context.Context, username string) (int, error)
	GetById(ctx context.Context, id int) (*datamodel.User, error)
	GetByUsername(ctx context.Context, username string) (*datamodel.User, error)
	FindAll(context.Context) ([]*datamodel.User, error)
}

type UserRepository interface {
	SafeUserRepository
	Insert(context.Context, *datamodel.User) error
}

func NewUserRepository(dbContext *db.DbContext) UserRepository {
	return &userRepository{
		dbContext: dbContext,
	}
}

type userRepository struct {
	dbContext *db.DbContext
}

func (ur userRepository) CountByUsername(ctx context.Context, username string) (int, error) {
	var count int64
	users := make([]*datamodel.User, 0)
	return int(count), ur.dbContext.GetDB(ctx).Find(&users, "username = ?", username).Count(&count).Error
}

func (ur userRepository) GetById(ctx context.Context, id int) (*datamodel.User, error) {
	user := new(datamodel.User)
	err := ur.dbContext.GetDB(ctx).First(user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, server_errors.NewNotFoundError("user")
	}
	return user, err
}

func (ur userRepository) GetByUsername(ctx context.Context, username string) (*datamodel.User, error) {
	user := new(datamodel.User)
	err := ur.dbContext.GetDB(ctx).First(user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, server_errors.NewNotFoundError("user")
	}
	return user, err
}

func (ur userRepository) FindAll(ctx context.Context) ([]*datamodel.User, error) {
	users := make([]*datamodel.User, 0)
	return users, ur.dbContext.GetDB(ctx).Find(&users).Error
}

func (ur userRepository) Insert(ctx context.Context, user *datamodel.User) error {
	return ur.dbContext.GetDB(ctx).Create(user).Error
}

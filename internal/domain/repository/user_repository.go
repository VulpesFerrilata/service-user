package repository

import (
	"context"
	"errors"

	"github.com/VulpesFerrilata/user/internal/domain/model"
	"gorm.io/gorm"

	"github.com/VulpesFerrilata/library/pkg/db"
	server_errors "github.com/VulpesFerrilata/library/pkg/errors"
)

type ReadOnlyUserRepository interface {
	CountByUsername(ctx context.Context, username string) (int, error)
	GetById(ctx context.Context, id int) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	FindAll(context.Context) ([]*model.User, error)
}

type UserRepository interface {
	ReadOnlyUserRepository

	Insert(context.Context, *model.User) error
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
	users := make([]*model.User, 0)
	return int(count), ur.dbContext.GetDB(ctx).Find(users, "user_name = ?", username).Count(&count).Error
}

func (ur userRepository) GetById(ctx context.Context, id int) (*model.User, error) {
	user := new(model.User)
	err := ur.dbContext.GetDB(ctx).First(user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, server_errors.NewNotFoundError("user")
	}
	return user, err
}

func (ur userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := new(model.User)
	err := ur.dbContext.GetDB(ctx).First(user, "username = ?", username).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, server_errors.NewNotFoundError("user")
	}
	return user, err
}

func (ur userRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	users := make([]*model.User, 0)
	return users, ur.dbContext.GetDB(ctx).Find(&users).Error
}

func (ur userRepository) Insert(ctx context.Context, user *model.User) error {
	return ur.dbContext.GetDB(ctx).Create(user).Error
}

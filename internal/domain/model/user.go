package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
)

type User struct {
	model.Model
	ID           int    `gorm:"primaryKey" validate:"required"`
	Username     string `gorm:"type:varchar(20);uniqueIndex" validate:"required"`
	HashPassword []byte `validate:"required"`
}

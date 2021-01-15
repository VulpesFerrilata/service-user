package model

import (
	"github.com/VulpesFerrilata/library/pkg/model"
	"github.com/google/uuid"
)

type User struct {
	model.Model
	ID           uuid.UUID `gorm:"type:uuid;primaryKey" validate:"required"`
	Username     string    `gorm:"type:varchar(20);uniqueIndex" validate:"required"`
	HashPassword []byte    `validate:"required"`
	DisplayName  string    `gorm:"type:varchar(50);unique" validate:"required"`
	Email        string    `gorm:"type:varchar(100);unique" validate:"required"`
}

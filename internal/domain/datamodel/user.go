package datamodel

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username     string `gorm:"type:varchar(20);uniqueIndex"`
	HashPassword []byte
}

package dto

import (
	"github.com/VulpesFerrilata/user/internal/domain/model"
)

func NewUserDTO(user *model.User) (*UserDTO, error) {
	userDTO := new(UserDTO)
	userDTO.ID = int(user.ID)
	userDTO.Username = user.Username
	return userDTO, nil
}

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

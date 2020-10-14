package response

import (
	"github.com/VulpesFerrilata/user/internal/domain/model"
)

func NewUserResponse(user *model.User) *UserResponse {
	userResponse := new(UserResponse)
	userResponse.ID = int(user.ID)
	userResponse.Username = user.Username
	return userResponse
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

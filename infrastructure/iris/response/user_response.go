package response

import "github.com/VulpesFerrilata/user/internal/usecase/dto"

func NewUserResponse(userDTO *dto.UserDTO) *UserResponse {
	userResponse := new(UserResponse)
	userResponse.ID = userDTO.ID
	userResponse.Username = userDTO.Username
	return userResponse
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

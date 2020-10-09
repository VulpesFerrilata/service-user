package viewmodel

import (
	"github.com/VulpesFerrilata/grpc/protoc/user"
	"github.com/VulpesFerrilata/user/internal/usecase/dto"
)

func NewUserResponse(userResponsePb *user.UserResponse) *UserResponse {
	return &UserResponse{
		userResponsePb: userResponsePb,
	}
}

type UserResponse struct {
	userResponsePb *user.UserResponse
}

func (ur *UserResponse) FromUserDTO(userDTO *dto.UserDTO) {
	ur.userResponsePb.ID = int64(userDTO.ID)
}

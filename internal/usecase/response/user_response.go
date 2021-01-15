package response

import "github.com/VulpesFerrilata/user/internal/domain/datamodel"

func NewUserResponse(user *datamodel.User) *UserResponse {
	userResponse := new(UserResponse)
	userResponse.ID = user.GetId().String()
	userResponse.Username = user.GetUsername()
	userResponse.DisplayName = user.GetDisplayName()
	userResponse.Email = user.GetEmail()
	return userResponse
}

type UserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
}

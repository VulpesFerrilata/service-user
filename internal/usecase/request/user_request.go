package request

type UserRequest struct {
	ID int `json:"id" validate:"required,gt=0"`
}

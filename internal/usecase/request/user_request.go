package request

type UserRequest struct {
	ID string `json:"id" validate:"required"`
}

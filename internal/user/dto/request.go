package dto


type CreateUserRequest struct {

	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required,min=11"`
	Password string `json:"password" validate:"required,min=8"`
}

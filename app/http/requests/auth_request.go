package requests

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

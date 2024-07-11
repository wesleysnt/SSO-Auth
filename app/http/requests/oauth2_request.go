package requests

type GrantType string

const (
	GrantTypeAuthCode           GrantType = "authorization_code"
	GrantTypePasswordCredential GrantType = "password"
	GrantTypeClientCredential   GrantType = "client_credentials"
)

type OAuth2Request struct {
	Email    string `validate:"required,email"`
	Name     string `validate:"required"`
	Password string `validate:"required"`
	Phone    string `validate:"required"`
}

type OAuth2LoginRequest struct {
	Email         string    `validate:"required,email" form:"email" json:"email"`
	Password      string    `validate:"required" form:"password" json:"password"`
	ClientId      string    `validate:"required" form:"client_id" json:"client_id"`
	GrantType     GrantType `validate:"required" form:"grant_type" json:"grant_type"`
	CodeChallenge string    `validate:"required_if=GrantType authorization_code" form:"code_challenge" json:"code_challenge"`
}

type IsLoggedInRequest struct {
	Token    string `validate:"required" form:"token" json:"token"`
	ClientId string `validate:"required" form:"client_id" json:"client_id"`
}

type RequestForgotPasswordRequest struct {
	Email string `validate:"required,email" form:"email" json:"email"`
}

type ResetPasswordRequest struct {
	Password        string `validate:"required" form:"password" json:"password"`
	ConfirmPassword string `validate:"required" form:"confirm_password" json:"confirm_password"`
	Token           string `validate:"required" form:"token" json:"token"`
}

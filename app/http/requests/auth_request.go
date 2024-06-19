package requests

type GrantType string

const (
	GrantTypeAuthCode           GrantType = "authorization_code"
	GrantTypePasswordCredential GrantType = "password"
	GrantTypeClientCredential   GrantType = "client_credentials"
)

type AuthRequest struct {
	Email    string `validate:"required,email"`
	Username string `validate:"required,alphanum,min=5,max=50"`
	Password string `validate:"required"`
}

type LoginRequest struct {
	Username string `validate:"required" form:"username" json:"username"`
	Password string `validate:"required" form:"password" json:"password"`
	ClientId uint   `validate:"required" form:"client_id" json:"client_id"`
}

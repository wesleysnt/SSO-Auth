package requests

type GrantType string

const (
	GrantTypeAuthCode           GrantType = "authorization_code"
	GrantTypePasswordCredential GrantType = "password"
	GrantTypeClientCredential   GrantType = "client_credentials"
)

type OAuth2Request struct {
	Email    string `validate:"required,email"`
	Name     string `validate:"required,alphanum,min=5,max=50"`
	Password string `validate:"required"`
	Phone    string `validate:"required"`
}

type OAuth2LoginRequest struct {
	Email       string `validate:"required" form:"email" json:"email"`
	Password    string `validate:"required" form:"password" json:"password"`
	ClientId    uint   `validate:"" form:"client_id" json:"client_id"`
	GrantType   string `validate:"required" form:"grant_type" json:"grant_type"`
	RedirectUri string `validate:"" form:"redirect_uri" json:"redirect_uri"`
}

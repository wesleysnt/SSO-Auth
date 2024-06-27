package requests

type TokenRequest struct {
	AuthCode  string `validate:"required" form:"auth_code" json:"auth_code"`
	ClientId  uint   `validate:"required" form:"client_id" json:"client_id"`
	UserId    uint   `validate:"required" form:"user_id" json:"user_id"`
	Scope     string `form:"scope" json:"scope"`
	GrantType string `validate:"required" form:"grant_type" json:"grant_type"`
}
<<<<<<< HEAD

type ValidateTokenRequest struct {
	Token  string `validate:"required" form:"token" json:"token"`
	Secret string `validate:"required" form:"secret" json:"secret"`
}

type RefreshTokenRequest struct {
	RefreshToken string `validate:"required" form:"refresh_token" json:"refresh_token"`
	Secret       string `validate:"required" form:"secret" json:"secret"`
	ClientId     uint   `validate:"required" form:"client_id" json:"client_id"`
	UserId       uint   `validate:"required" form:"user_id" json:"user_id"`
}
=======
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)

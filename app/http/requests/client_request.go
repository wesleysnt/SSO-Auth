package requests

type ClientRequest struct {
	ClientId    string `validate:"required" form:"client_id"`
	Secret      string `validate:"required" form:"secret"`
	RedirectUri string `validate:"required" form:"redirect_uri"`
}

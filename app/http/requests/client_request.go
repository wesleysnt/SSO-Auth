package requests

type ClientRequest struct {
	Name        string `validate:"required" form:"name"`
	Secret      string `validate:"required" form:"secret"`
	RedirectUri string `validate:"required" form:"redirect_uri"`
}

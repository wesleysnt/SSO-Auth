package responses

type ClientDetail struct {
	Id          uint   `json:"id"`
	ClientId    string `json:"client_id"`
	Secret      string `json:"secret"`
	RedirectUri string `json:"redirect_uri"`
}

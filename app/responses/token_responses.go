package responses

type TokenResponse struct {
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
	Scope        string       `json:"scope"`
}

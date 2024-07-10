package responses

import "github.com/golang-jwt/jwt/v5"

type TokenResponse struct {
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
	Scope        string       `json:"scope"`
	RedirectUri  string       `json:"redirect_uri"`
}

type ValidateTokenResponse struct {
	Active   bool            `json:"active"`
	Exp      jwt.NumericDate `json:"exp"`
	ClientId string          `json:"client_id"`
	UserId   uint            `json:"user_id"`
}

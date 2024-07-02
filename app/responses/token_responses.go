package responses

import "github.com/golang-jwt/jwt/v5"

type TokenResponse struct {
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
	Scope        string       `json:"scope"`
}

type ValidateTokenResponse struct {
	Active   bool            `json:"active"`
	Exp      jwt.NumericDate `json:"exp"`
	ClientId uint            `json:"client_id"`
	UserId   uint            `json:"user_id"`
}

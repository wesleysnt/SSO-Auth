package responses

import "github.com/golang-jwt/jwt/v5"

type LoginResponses struct {
	Id           uint         `json:"id"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
}

type AccessToken struct {
	Token      string           `json:"token"`
	ExpiryTime *jwt.NumericDate `json:"expiry_time"`
}

type RefreshToken struct {
	Token      string           `json:"token"`
	ExpiryTime *jwt.NumericDate `json:"expiry_time"`
}

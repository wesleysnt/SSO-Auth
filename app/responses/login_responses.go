package responses

import (
	"github.com/golang-jwt/jwt/v5"
)

type AdminLoginResponses struct {
	Id          uint        `json:"id"`
	Email       string      `json:"email"`
	AccessToken AccessToken `json:"access_token"`
}

type LoginResponses struct {
	Id           uint         `json:"id"`
	Username     string       `json:"username"`
	Email        string       `json:"email"`
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
	RedirectUri  string       `json:"redirect_uri"`
}

type LoginResponsesAuthCode struct {
	Id          uint     `json:"id"`
	Username    string   `json:"username"`
	Email       string   `json:"email"`
	AuthCode    AuthCode `json:"auth_code"`
	RedirectUri string   `json:"redirect_uri"`
}

type AuthCode struct {
	Code       string `json:"code"`
	ExpiryTime int64  `json:"expiry_time"`
}

type AccessToken struct {
	Token      string           `json:"token"`
	ExpiryTime *jwt.NumericDate `json:"expiry_time"`
}

type RefreshToken struct {
	Token      string           `json:"token"`
	ExpiryTime *jwt.NumericDate `json:"expiry_time"`
}

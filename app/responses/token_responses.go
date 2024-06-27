package responses

<<<<<<< HEAD
import "github.com/golang-jwt/jwt/v5"

=======
>>>>>>> 325f9fc (.)
type TokenResponse struct {
	AccessToken  AccessToken  `json:"access_token"`
	RefreshToken RefreshToken `json:"refresh_token"`
	Scope        string       `json:"scope"`
}
<<<<<<< HEAD

type ValidateTokenResponse struct {
	Active   bool            `json:"active"`
	Exp      jwt.NumericDate `json:"exp"`
	ClientId uint            `json:"client_id"`
	UserId   uint            `json:"user_id"`
}
=======
>>>>>>> 325f9fc (.)

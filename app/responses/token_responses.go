package responses

<<<<<<< HEAD
import "github.com/golang-jwt/jwt/v5"

=======
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
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
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)

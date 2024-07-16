package middleware

import (
	"os"
	"sso-auth/app/helpers"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewAuthMMiddleware() fiber.Handler {
	secret := os.Getenv("JWT_SECRET")
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(secret)},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err == fiber.ErrInternalServerError {
				return helpers.ResponseApiError(ctx, err.Error(), 500, nil)
			}
			return helpers.ResponseApiError(ctx, err.Error(), 500, nil)
		},
	})
}

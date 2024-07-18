package tokenhandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/http/requests"
	"sso-auth/app/schemas"
	"sso-auth/app/services"
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

type TokenHandler struct {
	tokenService *services.Oauth2TokenService
}

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{
		tokenService: services.NewOauth2TokenService(),
	}
}

func (h *TokenHandler) Token(c *fiber.Ctx) (err error) {
	data := requests.TokenRequest{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err = pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)

	}

	res, err := h.tokenService.Token(&data, c.UserContext())
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiCreated(c, "Login successful", res)
}

func (h *TokenHandler) ValidateToken(c *fiber.Ctx) (err error) {
	data := requests.ValidateTokenRequest{}

	c.BodyParser(&data)

	res, err := h.tokenService.ValidateToken(&data, c.UserContext())
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiOk(c, "Token is valid", res)
}

func (h *TokenHandler) RefreshToken(c *fiber.Ctx) (err error) {
	data := requests.RefreshTokenRequest{}

	c.BodyParser(&data)

	res, err := h.tokenService.RefreshToken(&data, c.UserContext())
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiCreated(c, "Token refreshed successfully", res)
}

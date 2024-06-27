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
	grantType := c.Query("grant_type", "")
	redirectUri := c.Query("redirect_uri", "")
	data := requests.TokenRequest{}

	c.BodyParser(&data)

	if grantType == "" {
		return helpers.ResponseApiBadRequest(c, "Grant type is required", nil)

	}

	if redirectUri == "" {
		return helpers.ResponseApiBadRequest(c, "Redirect uri is required", nil)

	}

	pkg.NewValidator()
	err = pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)

	}

	res, err := h.tokenService.Token(&data, grantType, redirectUri)
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiCreated(c, "Login successful", res)
}

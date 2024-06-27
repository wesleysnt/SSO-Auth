package authhandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/http/requests"
	"sso-auth/app/schemas"
	"sso-auth/app/services"
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

type OAuth2Handler struct {
	authService *services.OAuth2Service
}

func NewAuthHandler() *OAuth2Handler {
	return &OAuth2Handler{
		authService: services.NewOAuth2Service(),
	}
}

func (h *OAuth2Handler) Register(c *fiber.Ctx) error {
	data := requests.OAuth2Request{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	err = h.authService.Register(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "User created", nil)
}

func (h *OAuth2Handler) Login(c *fiber.Ctx) (err error) {
	data := requests.OAuth2LoginRequest{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err = pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)

	}

	res, err := h.authService.Login(&data)
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	c.Request().Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return helpers.ResponseApiCreated(c, "Login successful", res)
}

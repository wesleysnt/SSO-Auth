package authhandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/http/requests"
	"sso-auth/app/schemas"
	"sso-auth/app/services"
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: services.NewAuthService(),
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	data := requests.LoginRequest{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	res, errLogin := h.authService.Login(&data)

	if errLogin != nil {
		respErr := errLogin.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiCreated(c, "successful", res)
}

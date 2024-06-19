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

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	data := requests.AuthRequest{}

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

func (h *AuthHandler) Login(c *fiber.Ctx) (err error) {
	grantType := c.Query("grant_type", "")
	redirectUri := c.Query("redirect_uri", "")
	data := requests.LoginRequest{}

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

	res, err := h.authService.Login(grantType, redirectUri, &data)
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}
	return helpers.ResponseApiCreated(c, "Successfully to create data", res)
}

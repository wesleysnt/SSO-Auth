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

	res, err := h.authService.Register(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "Registration success, please check your email to verify account", res)
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
	return helpers.ResponseApiCreated(c, "Login successful", res)
}

func (h *OAuth2Handler) VerifOtp(c *fiber.Ctx) error {
	data := requests.VerifOtp{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	res, err := h.authService.VerifOtp(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "Verification success", res)
}

func (h *OAuth2Handler) IsLoggedIn(c *fiber.Ctx) error {
	data := requests.IsLoggedInRequest{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	res, err := h.authService.IsLoggedIn(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "User had logged in!", res)
}

func (h *OAuth2Handler) RequestForgotPassword(c *fiber.Ctx) error {
	data := requests.RequestForgotPasswordRequest{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	err = h.authService.RequestForgotPassword(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "Please check your email to reset your password", nil)
}

func (h *OAuth2Handler) ResetPassword(c *fiber.Ctx) error {
	data := requests.ResetPasswordRequest{}

	c.BodyParser(&data)

	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	if data.Password != data.ConfirmPassword {
		return helpers.ResponseApiBadRequest(c, "Password and password confirmation must be same", nil)
	}

	err = h.authService.ResetPassword(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)

		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "Password has been reset", nil)
}

package userclientloghandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/schemas"
	"sso-auth/app/services"

	"github.com/gofiber/fiber/v2"
)

type UserClientLogHandler struct {
	userClientLogService *services.UserClientLogService
}

func NewUserClientLogHandler() *UserClientLogHandler {
	return &UserClientLogHandler{
		userClientLogService: services.NewUserClientLogService(),
	}
}

func (h *UserClientLogHandler) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	sort := c.Query("sort")

	res, err := h.userClientLogService.List(page, limit, sort, c.UserContext())
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)
}

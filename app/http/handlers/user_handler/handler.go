package userhandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/schemas"
	"sso-auth/app/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userService: services.NewUserService(),
	}
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	sort := c.Query("sort")

	res, err := h.userService.List(page, limit, sort, c.UserContext())
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)
}

func (h *UserHandler) GetById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id", 0)

	res, err := h.userService.GetById(uint(id), c.UserContext())
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)
}

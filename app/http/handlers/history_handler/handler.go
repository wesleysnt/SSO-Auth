package historyhandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/schemas"
	"sso-auth/app/services"

	"github.com/gofiber/fiber/v2"
)

type HistoryHandler struct {
	historyService *services.HistoryService
}

func NewHistoryHandler() *HistoryHandler {
	return &HistoryHandler{
		historyService: services.NewHistoryService(),
	}
}

func (h *HistoryHandler) TokenHistory(c *fiber.Ctx) error {

	page := c.Query("page")
	limit := c.Query("limit")
	sort := c.Query("sort")

	res, err := h.historyService.TokenHistory(page, limit, sort, c.UserContext())

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)

}

func (h *HistoryHandler) RefreshTokenHistory(c *fiber.Ctx) error {

	page := c.Query("page")
	limit := c.Query("limit")
	sort := c.Query("sort")

	res, err := h.historyService.RefreshTokenHistory(page, limit, sort, c.UserContext())

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)

}

func (h *HistoryHandler) AuthCodeHistory(c *fiber.Ctx) error {

	page := c.Query("page")
	limit := c.Query("limit")
	sort := c.Query("sort")

	res, err := h.historyService.AuthCodeHistory(page, limit, sort, c.UserContext())

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)

}

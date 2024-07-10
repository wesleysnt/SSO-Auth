package dashboardhandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/services"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	dashboardService *services.DashboardService
}

func NewDashboardHandler() *DashboardHandler {
	return &DashboardHandler{
		dashboardService: services.NewDashboardService(),
	}
}

func (h *DashboardHandler) Get(c *fiber.Ctx) error {
	res, _ := h.dashboardService.Get()
	if res == nil {

		return helpers.ResponseApiError(c, "Something wrong while getting dashboard data", 500, nil)
	}

	return helpers.ResponseApiOk(c, "OK", res)
}

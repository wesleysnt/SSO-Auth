package dashboardhandler

import (
	"sso-auth/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func DashboardRoute(route fiber.Router) {

	handler := NewDashboardHandler()
	auth := route.Group("/dashboard")

	auth.Get("/", middleware.NewAuthMMiddleware(), handler.Get)
}

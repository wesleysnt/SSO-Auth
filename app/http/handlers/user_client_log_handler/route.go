package userclientloghandler

import (
	"sso-auth/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserClientLogRoute(route fiber.Router) {

	handler := NewUserClientLogHandler()
	auth := route.Group("/logs")

	auth.Get("/", middleware.NewAuthMMiddleware(), handler.List)
}

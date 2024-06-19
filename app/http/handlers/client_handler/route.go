package clienthandler

import (
	"sso-auth/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func ClientRoute(route fiber.Router) {

	handler := NewClientHandler()
	auth := route.Group("/client")

	auth.Post("/", middleware.NewAuthMMiddleware(), handler.Create)
}

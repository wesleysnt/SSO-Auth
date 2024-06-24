package authhandler

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(route fiber.Router) {

	handler := NewAuthHandler()
	auth := route.Group("/auth")

	auth.Post("/login", handler.Login)
}

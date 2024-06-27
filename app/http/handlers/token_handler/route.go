package tokenhandler

import "github.com/gofiber/fiber/v2"

func TokenRoute(route fiber.Router) {

	handler := NewTokenHandler()
	auth := route.Group("/oauth2")

	auth.Post("/token", handler.Token)
}

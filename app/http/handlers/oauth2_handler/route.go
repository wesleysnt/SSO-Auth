package authhandler

import "github.com/gofiber/fiber/v2"

func AuthRoute(route fiber.Router) {

	handler := NewAuthHandler()
	auth := route.Group("/oauth2")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
}

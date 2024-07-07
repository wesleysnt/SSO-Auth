package authhandler

import "github.com/gofiber/fiber/v2"

func OAuth2Route(route fiber.Router) {

	handler := NewAuthHandler()
	auth := route.Group("/oauth2")

	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)
	auth.Post("/verif-otp", handler.VerifOtp)
}

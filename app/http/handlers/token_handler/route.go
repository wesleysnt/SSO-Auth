package tokenhandler

import "github.com/gofiber/fiber/v2"

func TokenRoute(route fiber.Router) {

	handler := NewTokenHandler()
	auth := route.Group("/oauth2")

	auth.Post("/token", handler.Token)
<<<<<<< HEAD
	auth.Post("/validate-token", handler.ValidateToken)
	auth.Post("/refresh-token", handler.RefreshToken)
=======
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
}

package clienthandler

import (
	"sso-auth/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func ClientRoute(route fiber.Router) {

	handler := NewClientHandler()
	auth := route.Group("/client")

	auth.Post("/", middleware.NewAuthMMiddleware(), handler.Create)
	auth.Get("/", middleware.NewAuthMMiddleware(), handler.List)
	auth.Get("/:id", middleware.NewAuthMMiddleware(), handler.Detail)
	auth.Put("/:id", middleware.NewAuthMMiddleware(), handler.Update)
	auth.Get("/generate-secret/:id", middleware.NewAuthMMiddleware(), handler.GenerateSecret)
	auth.Delete("/:id", middleware.NewAuthMMiddleware(), handler.Delete)
}

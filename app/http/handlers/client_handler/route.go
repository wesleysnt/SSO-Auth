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
	auth.Get("/generate-secret", middleware.NewAuthMMiddleware(), handler.GenerateSecret)
	auth.Get("/:id", middleware.NewAuthMMiddleware(), handler.Detail)
	auth.Put("/:id", middleware.NewAuthMMiddleware(), handler.Update)
	auth.Delete("/:id", middleware.NewAuthMMiddleware(), handler.Delete)
}

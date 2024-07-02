package userhandler

import (
	"sso-auth/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(route fiber.Router) {

	handler := NewUserHandler()
	auth := route.Group("/user")

	auth.Get("/", middleware.NewAuthMMiddleware(), handler.List)
	auth.Get("/:id", middleware.NewAuthMMiddleware(), handler.GetById)
}

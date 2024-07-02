package historyhandler

import (
	"sso-auth/app/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func HistoryRoute(route fiber.Router) {

	handler := NewHistoryHandler()
	auth := route.Group("/history")

	auth.Get("/token", middleware.NewAuthMMiddleware(), handler.TokenHistory)
	auth.Get("/refresh-token", middleware.NewAuthMMiddleware(), handler.RefreshTokenHistory)
	auth.Get(("/auth-code"), middleware.NewAuthMMiddleware(), handler.AuthCodeHistory)
}

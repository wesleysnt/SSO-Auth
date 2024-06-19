package routes

import (
	clienthandler "sso-auth/app/http/handlers/client_handler"
	authhandler "sso-auth/app/http/handlers/oauth2_handler"
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	authhandler.AuthRoute(api)

	clienthandler.ClientRoute(api)
	pkg.ListRouters(app)

}

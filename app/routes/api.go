package routes

import (
	authhandler "sso-auth/app/http/handlers/auth_handler"
	clienthandler "sso-auth/app/http/handlers/client_handler"
	oauth2handler "sso-auth/app/http/handlers/oauth2_handler"
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	oauth2handler.OAuth2Route(api)

	authhandler.AuthRoute(api)

	clienthandler.ClientRoute(api)
	pkg.ListRouters(app)

}

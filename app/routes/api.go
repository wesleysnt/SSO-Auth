package routes

import (
	authhandler "sso-auth/app/http/handlers/auth_handler"
	clienthandler "sso-auth/app/http/handlers/client_handler"
	historyhandler "sso-auth/app/http/handlers/history_handler"
	oauth2handler "sso-auth/app/http/handlers/oauth2_handler"
	tokenhandler "sso-auth/app/http/handlers/token_handler"
<<<<<<< HEAD
	userhandler "sso-auth/app/http/handlers/user_handler"
=======
<<<<<<< HEAD
>>>>>>> 325f9fc (.)
=======
>>>>>>> ec32a2f (.)
>>>>>>> 7a7a01f (.)
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	oauth2handler.OAuth2Route(api)

	tokenhandler.TokenRoute(api)

	authhandler.AuthRoute(api)

	historyhandler.HistoryRoute(api)

	clienthandler.ClientRoute(api)

	userhandler.UserRoute(api)

	pkg.ListRouters(app)

}

package main

import (
	"log"
	"sso-auth/app/configs"
	"sso-auth/app/helpers"
	"sso-auth/app/routes"
	"sso-auth/app/schemas"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := configs.InitEnv()

	if err != nil {
		panic("Cannot load env!")
	}

	configApp := configs.LoadAppConfig()
	app := fiber.New(fiber.Config{
		CaseSensitive:         true,
		StrictRouting:         true,
		AppName:               configApp.AppName,
		DisableStartupMessage: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			respErr := err.(*schemas.ResponseApiError)
			catchErr := helpers.CatchErrorResponseApi(respErr)

			if err == fiber.ErrInternalServerError {
				return helpers.ResponseApiError(ctx, catchErr.Message, catchErr.StatusCode, nil)
			}
			return helpers.ResponseApiError(ctx, catchErr.Message, catchErr.StatusCode, nil)
		},
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	configs.ConnectDB()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":" + configApp.AppPort))

}

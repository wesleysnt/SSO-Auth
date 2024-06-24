package main

import (
	"log"
	"sso-auth/app/configs"
	"sso-auth/app/database/seeders"
	"sso-auth/app/helpers"
	"sso-auth/app/routes"
	"sso-auth/cmd/cli"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	err := configs.InitEnv()

	if err != nil {
		panic("Cannot load env!")
	}

	configApp := configs.LoadAppConfig()
	cli.Execute()
	app := fiber.New(fiber.Config{
		CaseSensitive:         true,
		StrictRouting:         true,
		AppName:               configApp.AppName,
		DisableStartupMessage: true,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if err == fiber.ErrInternalServerError {
				return helpers.ResponseApiError(ctx, err.Error(), 500, nil)
			}
			return helpers.ResponseApiError(ctx, err.Error(), 500, nil)
		},
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	configs.ConnectDB()

	seeders.RegisterSeeder()
	routes.RegisterRoutes(app)
	log.Fatal(app.Listen(":" + configApp.AppPort))

}

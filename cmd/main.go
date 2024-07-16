package main

import (
	"log"
	"os"
	"sso-auth/app/configs"
	"sso-auth/app/helpers"
	"sso-auth/app/routes"
	"sso-auth/cmd/cli/commands"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gookit/color"
)

func main() {
	err := configs.InitEnv()

	if err != nil {
		panic("Cannot load env!")
	}

	configApp := configs.LoadAppConfig()

	if len(os.Args) >= 2 {
		configs.ConnectDB()
		commands.Execute()
		return
	}

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

	routes.RegisterRoutes(app)
	color.Blueln("This app is running on 127.0.0.1:" + configApp.AppPort)
	log.Fatal(app.Listen(":" + configApp.AppPort))

}

package configs

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ConfigApp struct {
	AppName  string
	AppPort  string
	AppDebug bool
	AppEnv   string
}

func InitEnv() error {
	err := godotenv.Load("./.env")

	if err != nil {
		return err
	}

	return err

}

func LoadAppConfig() (config ConfigApp) {
	appDebug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	config = ConfigApp{
		AppName:  os.Getenv("APP_NAME"),
		AppPort:  os.Getenv("APP_PORT"),
		AppDebug: appDebug,
		AppEnv:   os.Getenv("APP_ENV"),
	}
	return
}

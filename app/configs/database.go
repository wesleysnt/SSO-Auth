package configs

import (
	"fmt"
	"os"
	"sso-auth/app/facades"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dbCon struct {
	username, password, db, host, port string
}

var (
	dbInstance *gorm.DB
	err        error
)

func InitDB() (db *gorm.DB, err error) {
	dbCon := dbCon{
		username: os.Getenv("DB_USERNAME"),
		password: os.Getenv("DB_PASSWORD"),
		db:       os.Getenv("DB_DATABASE"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbCon.host, dbCon.username, dbCon.password, dbCon.db, dbCon.port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return

}

func ConnectDB() {
	if dbInstance == nil {
		dbInstance, err = InitDB()
	}

	if err != nil {
		panic("Failed to connect to database")
	}
	facades.MakeOrm(dbInstance)
}

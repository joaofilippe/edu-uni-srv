package appconfig

import (
	"errors"
	"os"

	"github.com/joaofilippe/edu-uni-srv/utils/logger"
	"github.com/joho/godotenv"
)

type App struct {
	ConfigPath string
	SecretKey  string
	Port       string
}

func NewApp(log *logger.Logger) *App {
	var app App

	app.ConfigPath = "./config"
	err := godotenv.Load(app.ConfigPath + "/.env")
	if err != nil {
		log.Fatalf(errors.New("can't load .env file :( "))
	}

	app.SecretKey = os.Getenv("SECRET_KEY")
	app.Port = os.Getenv("PORT")

	return &app
}

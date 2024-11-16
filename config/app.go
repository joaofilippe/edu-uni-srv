package appconfig

import (
	"errors"
	"os"
	"sync"

	"github.com/joaofilippe/edu-uni-srv/common/logger"

	"github.com/joho/godotenv"
)

var (
	instance *App
	once     sync.Once
)

type App struct {
	ConfigPath string
	SecretKey  string
	Port       string
}

func NewApp(log *logger.Logger) *App {
	once.Do(func() {
		var app App

		app.ConfigPath = "./config"
		err := godotenv.Load(app.ConfigPath + "/.env")
		if err != nil {
			log.Fatalf(errors.New("can't load .env file :( "))
		}

		app.SecretKey = os.Getenv("SECRET_KEY")
		app.Port = os.Getenv("PORT")

		instance = &app
	})

	return instance

}

func Instance() *App {
	return instance
}
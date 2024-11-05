package main

import (
	"fmt"

	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/joaofilippe/edu-uni-srv/config"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	"github.com/joaofilippe/edu-uni-srv/infra/server"
	"github.com/joaofilippe/edu-uni-srv/utils/logger"
)

func main() {
	logger := logger.NewLogger()
	appConfig := config.NewApp(logger)

	conn := getDbConnection(logger, appConfig)

	fmt.Println(conn.Config.Dsn)
	
	app := application.Application{}

	server := server.NewServer(&app)

	server.Start()
}

func getDbConnection(log *logger.Logger, appConfig *config.App) *database.Connection {
		return database.GetConfigFromYaml(log, appConfig)
}
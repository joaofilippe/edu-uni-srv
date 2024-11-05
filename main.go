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
	conn.Connection.Exec(table)

	fmt.Println(conn.ConfigDB.Dsn)

	app := application.Application{}

	server := server.NewServer(&app)

	server.Start()
}

func getDbConnection(log *logger.Logger, appConfig *config.App) *database.Connection {
	return database.GetConfigFromYaml(log, appConfig)
}

const table = `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
`

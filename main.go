package main

import (
	"github.com/joaofilippe/edu-uni-srv/config"
	"github.com/joaofilippe/edu-uni-srv/di"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	"github.com/joaofilippe/edu-uni-srv/infra/database/migrations"
	"github.com/joaofilippe/edu-uni-srv/infra/server"
	"github.com/joaofilippe/edu-uni-srv/utils/logger"
)

func main() {
	l := logger.NewLogger()
	appConfig := appconfig.NewApp(l)

	conn := database.NewConnection(l, appConfig)
	application := di.ApplicationFactory(conn)

	err := conn.DBConnection.Ping()
	if err != nil {
		l.Fatalf(err)
	}

	err = migrations.RunMigrations(conn)
	if err != nil {
		l.Fatalf(err)
	}

	s := server.NewServer(application)

	err = s.Start()
	if err != nil {
		l.Fatalf(err)
	}
}

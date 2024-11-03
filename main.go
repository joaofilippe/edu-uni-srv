package main

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/joaofilippe/edu-uni-srv/infra/server"
)

func main() {
	app := application.Application{}

	server := server.NewServer(&app)

	server.Start()
}
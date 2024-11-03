package main

import (
	"github.com/joaofilippe/edu-uni-srv/core"
	"github.com/joaofilippe/edu-uni-srv/infra/server"
)

func main() {
	app := core.Application{}

	server := server.NewServer(&app)

	server.Start()
}
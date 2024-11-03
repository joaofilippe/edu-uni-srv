package server

import (
	"github.com/joaofilippe/edu-uni-srv/core"
	"github.com/labstack/echo/v4"
)

type Server struct {
	application *core.Application
	server      *echo.Echo
}

func NewServer(app *core.Application) *Server {
	e := echo.New()

	return &Server{
		application: app,
		server:      e,
	}
}

func (s *Server) Start() error {
	return s.server.Start(":8080")
}

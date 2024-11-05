package server

import (
	"github.com/joaofilippe/edu-uni-srv/api"
	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/labstack/echo/v4"
)

type Server struct {
	application *application.Application
	api         *api.Api
	server      *echo.Echo
}

func NewServer(app *application.Application) *Server {
	e := echo.New()

	return &Server{
		application: app,
		server:      e,
	}
}

func (s *Server) Start() error {
	err := s.registerRoutes()
	if err != nil {
		return err
	}
	return s.server.Start("127.0.0.1:5000")
}

func (s *Server) registerRoutes() error {
	s.api.BuildRoutes(s.server)
	return nil
}

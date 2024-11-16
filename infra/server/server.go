package server

import (
	"os"

	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/joaofilippe/edu-uni-srv/application/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	application *application.Application
	api         *api.Api
	server      *echo.Echo
}

func NewServer(app *application.Application) *Server {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	a := api.NewApi(app)

	return &Server{
		application: app,
		api:         a,
		server:      e,
	}
}

func (s *Server) Start() error {
	err := s.registerRoutes()
	if err != nil {
		return err
	}

	runningInDocker := os.Getenv("ENV")
	if runningInDocker == "docker" {
		return s.server.Start(":5000")
	} else {

		return s.server.Start("127.0.0.1:5000")
	}
}

func (s *Server) registerRoutes() error {
	s.api.BuildRoutes(s.server)
	return nil
}

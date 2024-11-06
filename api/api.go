package api

import (
	userweb "github.com/joaofilippe/edu-uni-srv/api/user"
	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/labstack/echo/v4"
)

type Api struct {
	userApi *userweb.UserWeb
}

func NewApi(application *application.Application) *Api {
	userService := application.UserService()
	return &Api{
		userApi: userweb.NewUserWeb(&userService),
	}
}

func (a *Api) BuildRoutes(server *echo.Echo) {
	a.userApi.BuildRoutes(server)
	a.registerWelcome(server)
}

func (a *Api) registerWelcome(server *echo.Echo) {
	server.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the API")
	})
}

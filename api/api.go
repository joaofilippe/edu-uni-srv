package api

import (
	userweb "github.com/joaofilippe/edu-uni-srv/api/user"
	"github.com/labstack/echo/v4"
)

type Api struct {
	userApi *userweb.UserWeb
}

func (a *Api) BuildRoutes(server *echo.Echo) {
	//a.userApi.BuildRoutes(server)
	a.registerWelcome(server)
}

func (a *Api) registerWelcome(server *echo.Echo) {
	server.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the API")
	})
}

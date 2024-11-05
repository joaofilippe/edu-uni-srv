package userweb

import "github.com/labstack/echo/v4"

func (uw *UserWeb) BuildRoutes(server *echo.Echo) {
	server.POST("/user", uw.CreateUser)
}

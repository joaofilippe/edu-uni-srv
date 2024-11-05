package userweb

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	"github.com/labstack/echo/v4"
)

type UserWeb struct {
	userService services.UserService
}

func (uw *UserWeb) CreateUser(c echo.Context) error {

	return c.JSON(200, "User created")
}

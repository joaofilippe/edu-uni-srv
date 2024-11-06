package userweb

import (
	"github.com/joaofilippe/edu-uni-srv/core/services"
	"github.com/labstack/echo/v4"
)

type UserWeb struct {
	userService iservices.IUserService
}

func NewUserWeb(userService *iservices.IUserService) *UserWeb {
	return &UserWeb{
		userService: *userService,
	}
}

func (uw *UserWeb) CreateUser(c echo.Context) error {

	return c.JSON(200, "User created")
}

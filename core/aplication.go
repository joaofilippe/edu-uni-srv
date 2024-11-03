package core

import "github.com/joaofilippe/edu-uni-srv/core/services"

type Application struct {
	UserService services.IUserService
}
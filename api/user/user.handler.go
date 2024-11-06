package userweb

import (
	"fmt"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/services"
	"github.com/labstack/echo/v4"
)

type UserWeb struct {
	userService iservices.IUserService
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	UserType string `json:"type"`
}

func NewUserWeb(userService *iservices.IUserService) *UserWeb {
	return &UserWeb{
		userService: *userService,
	}
}

func (uw *UserWeb) CreateUser(c echo.Context) error {
	req := new(CreateUserRequest)
	err := c.Bind(req)
	{
		if err != nil {
			return c.JSON(400, "Invalid request")
		}

		userType, err := enums.ParseUserType(req.UserType)
		if err != nil {
			return err
		}

		createUser := userEntities.NewCreateUser(req.Username, req.Password, req.Email, userType, nil)

		id, err := uw.userService.Create(createUser)
		if err != nil {
			return err
		}

		fmt.Println(id)

		return c.JSON(200, "User created")
	}
}

package userweb

import (
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	userService iservices.IUserService
}

func NewUserWeb(userService *iservices.IUserService) *WebService {
	return &WebService{
		userService: *userService,
	}
}

func (uw *WebService) CreateUser(c echo.Context) error {
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

		userID, err := uw.userService.Create(createUser)
		if err != nil {
			return err
		}

		adminentities.NewCreate(userID)

		return c.JSON(200, "User created")
	}
}

func (uw *WebService) Login(c echo.Context) error {
	req := new(LoginRequest)

	err := c.Bind(req)
	if err != nil {
		return c.JSON(400, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{
			"Invalid request",
			err.Error(),
		})
	}

	token, err := uw.userService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(400, struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}{
			"Invalid credentials",
			err.Error(),
		})
	}

	return c.JSON(200, struct {
		Message string `json:"message"`
		Token string `json:"token"`
	}{
		"Login successful",
		token,
	})

}

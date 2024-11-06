package adminweb

import (
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	iservices "github.com/joaofilippe/edu-uni-srv/core/services"
	"github.com/labstack/echo/v4"
)

type WebService struct {
	adminService iservices.IAdminService
	userService  iservices.IUserService
}

func NewWebService(
	adminService *iservices.IAdminService,
	userService *iservices.IUserService,
) *WebService {
	return &WebService{
		adminService: *adminService,
		userService:  *userService,
	}
}

func (aw *WebService) CreateAdmin(c echo.Context) error {
	req := new(CreateAdminRequest)
	err := c.Bind(req)

	if err != nil {
		return c.JSON(400, "Invalid request")
	}

	createUser := userEntities.NewCreateUser(
		req.Username,
		req.Password,
		req.Email,
		enums.Administrator,
		nil,
	)

	userID, err := aw.userService.Create(createUser)
	if err != nil {
		return err
	}

	_, err = aw.adminService.Create(adminentities.NewCreate(userID))
	if err != nil {
		return err
	}

	return c.JSON(200, "Admin created")

}

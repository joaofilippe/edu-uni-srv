package api

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/joaofilippe/edu-uni-srv/application/api/admin"
	studentweb "github.com/joaofilippe/edu-uni-srv/application/api/student"
	"github.com/joaofilippe/edu-uni-srv/application/api/user"
	"github.com/labstack/echo/v4"
)

type Api struct {
	userWebService  *userweb.WebService
	adminWebService *adminweb.WebService
	studentWebService *studentweb.WebService
}

func NewApi(application *application.Application) *Api {
	userService := application.UserService()
	adminService := application.AdminService()
	studentService := application.StudentService()
	return &Api{
		userWebService:  userweb.NewUserWeb(&userService),
		adminWebService: adminweb.NewWebService(&adminService, &userService),
		studentWebService: studentweb.NewStudentWeb(&studentService, &userService),
	}
}

func (a *Api) BuildRoutes(server *echo.Echo) {
	a.userWebService.BuildRoutes(server)
	a.adminWebService.BuildRoutes(server)
	a.studentWebService.BuildRoutes(server)
	a.registerWelcome(server)
}

func (a *Api) registerWelcome(server *echo.Echo) {
	server.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the API")
	})
}

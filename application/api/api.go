package api

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	"github.com/joaofilippe/edu-uni-srv/application/api/admin"
	"github.com/joaofilippe/edu-uni-srv/application/api/class"
	contentweb "github.com/joaofilippe/edu-uni-srv/application/api/content"
	guardianweb "github.com/joaofilippe/edu-uni-srv/application/api/guardian"
	studentweb "github.com/joaofilippe/edu-uni-srv/application/api/student"
	teacherweb "github.com/joaofilippe/edu-uni-srv/application/api/teacher"
	"github.com/joaofilippe/edu-uni-srv/application/api/user"
	"github.com/labstack/echo/v4"
)

type Api struct {
	userWebService     *userweb.WebService
	adminWebService    *adminweb.WebService
	studentWebService  *studentweb.WebService
	teacherWebService  *teacherweb.WebService
	guardianWebService *guardianweb.WebService
	contentWebService  *contentweb.WebService
	classWebService    *classweb.WebService
}

func NewApi(application *application.Application) *Api {
	userService := application.UserService()
	adminService := application.AdminService()
	studentService := application.StudentService()
	teacherService := application.TeacherService()
	guardianService := application.GuardianService()
	contentService := application.ContentService()
	classService := application.ClassService()
	return &Api{
		userweb.NewUserWeb(&userService),
		adminweb.NewWebService(&adminService, &userService),
		studentweb.NewStudentWeb(&studentService, &userService),
		teacherweb.NewTeacherWeb(&teacherService, &userService),
		guardianweb.NewGuardianWeb(&guardianService, &userService),
		contentweb.NewContentWeb(&contentService),
		classweb.NewClassWeb(&classService),
	}
}

func (a *Api) BuildRoutes(server *echo.Echo) {
	a.userWebService.BuildRoutes(server)
	a.adminWebService.BuildRoutes(server)
	a.studentWebService.BuildRoutes(server)
	a.teacherWebService.BuildRoutes(server)
	a.guardianWebService.BuildRoutes(server)
	a.contentWebService.BuildRoutes(server)
	a.classWebService.BuildRoutes(server)
	a.registerWelcome(server)
}

func (a *Api) registerWelcome(server *echo.Echo) {
	server.GET("/welcome", func(c echo.Context) error {
		return c.String(200, "Welcome to the API")
	})
}

package di

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	servicesdi "github.com/joaofilippe/edu-uni-srv/di/services"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	"github.com/joaofilippe/edu-uni-srv/repositories"
)

func ApplicationFactory(connection *database.DBConnection) *application.Application {
	a := application.Application{}
	adminRepository := repositories.NewAdminRepository(connection)
	teacherRepository := repositories.NewTeacherRepository(connection)
	studentRepository := repositories.NewStudentRepository(connection)
	guardianRepository := repositories.NewGuardianRepository(connection)
	userRepository := repositories.NewUserRepository(connection)

	userService := servicesdi.UserServiceFactory(userRepository, adminRepository, teacherRepository, studentRepository, guardianRepository)
	teacherService := servicesdi.TeacherServiceFactory(teacherRepository, userRepository)
	a.SetUserService(userService)
	return &a
}

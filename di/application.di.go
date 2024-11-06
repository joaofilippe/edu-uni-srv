package di

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	servicesdi "github.com/joaofilippe/edu-uni-srv/di/services"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	"github.com/joaofilippe/edu-uni-srv/repositories"
)

func ApplicationFactory(connection *database.DBConnection) *application.Application {
	userRepository := repositories.NewUserRepository(connection)
	adminRepository := repositories.NewAdminRepository(connection)
	teacherRepository := repositories.NewTeacherRepository(connection)
	studentRepository := repositories.NewStudentRepository(connection)
	guardianRepository := repositories.NewGuardianRepository(connection)

	userService := servicesdi.UserServiceFactory(userRepository, adminRepository, teacherRepository, studentRepository, guardianRepository)
	studentService := servicesdi.StudentServiceFactory(studentRepository, userRepository)
	teacherService := servicesdi.TeacherServiceFactory(teacherRepository, userRepository)
	guardianService := servicesdi.GuardianServiceFactory(guardianRepository, userRepository)
	adminService := servicesdi.AdminServiceFactory(adminRepository, userRepository)

	return application.NewApplication(
		userService,
		adminService,
		teacherService,
		studentService,
		guardianService,
	)
}

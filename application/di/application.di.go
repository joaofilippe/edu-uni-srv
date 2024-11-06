package di

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	servicesdi2 "github.com/joaofilippe/edu-uni-srv/application/di/services"
	repositories2 "github.com/joaofilippe/edu-uni-srv/application/repositories"
	"github.com/joaofilippe/edu-uni-srv/application/repositories/adminrepository"
	"github.com/joaofilippe/edu-uni-srv/application/repositories/userrepository"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

func ApplicationFactory(connection *database.DBConnection) *application.Application {
	userRepository := userrepository.NewUserRepository(connection)
	adminRepository := adminrepository.NewAdminRepository(connection)
	teacherRepository := repositories2.NewTeacherRepository(connection)
	studentRepository := repositories2.NewStudentRepository(connection)
	guardianRepository := repositories2.NewGuardianRepository(connection)

	userService := servicesdi2.UserServiceFactory(userRepository, adminRepository, teacherRepository, studentRepository, guardianRepository)
	studentService := servicesdi2.StudentServiceFactory(studentRepository, userRepository)
	teacherService := servicesdi2.TeacherServiceFactory(teacherRepository, userRepository)
	guardianService := servicesdi2.GuardianServiceFactory(guardianRepository, userRepository)
	adminService := servicesdi2.AdminServiceFactory(adminRepository, userRepository)

	return application.NewApplication(
		userService,
		adminService,
		teacherService,
		studentService,
		guardianService,
	)
}

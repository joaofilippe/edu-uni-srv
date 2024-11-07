package di

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	servicesdi "github.com/joaofilippe/edu-uni-srv/application/di/services"
	adminrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/admin_repository"
	guardianrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/guardian_repository"
	studentrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/student_repository"
	teacherrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/teacher_repository"
	"github.com/joaofilippe/edu-uni-srv/application/repositories/user_repository"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

func ApplicationFactory(connection *database.DBConnection) *application.Application {
	userRepository := userrepository.NewUserRepository(connection)
	adminRepository := adminrepository.NewAdminRepository(connection)
	teacherRepository := teacherrepository.NewTeacherRepository(connection)
	studentRepository := studentrepository.NewStudentRepository(connection)
	guardianRepository := guardianrepository.NewGuardianRepository(connection)

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

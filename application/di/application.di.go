package di

import (
	"github.com/joaofilippe/edu-uni-srv/application"
	servicesdi "github.com/joaofilippe/edu-uni-srv/application/di/services"
	adminrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/admin_repository"
	"github.com/joaofilippe/edu-uni-srv/application/repositories/class_repository"
	contentrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/content_repository"
	guardianrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/guardian_repository"
	studentrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/student_repository"
	teacherrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/teacher_repository"
	userrepository "github.com/joaofilippe/edu-uni-srv/application/repositories/user_repository"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

func ApplicationFactory(connection *database.DBConnection) *application.Application {
	userRepository := userrepository.NewUserRepository(connection)
	adminRepository := adminrepository.NewAdminRepository(connection)
	teacherRepository := teacherrepository.NewTeacherRepository(connection)
	studentRepository := studentrepository.NewStudentRepository(connection)
	guardianRepository := guardianrepository.NewGuardianRepository(connection)
	contentRepository := contentrepository.NewContentRepository(connection)
	classrepository := classrepository.NewClassRepository(connection)

	userService := servicesdi.UserServiceFactory(userRepository, adminRepository, teacherRepository, studentRepository, guardianRepository)
	studentService := servicesdi.StudentServiceFactory(studentRepository, userRepository)
	teacherService := servicesdi.TeacherServiceFactory(teacherRepository, userRepository)
	guardianService := servicesdi.GuardianServiceFactory(guardianRepository, studentRepository, userRepository)
	adminService := servicesdi.AdminServiceFactory(adminRepository, userRepository)
	contentService := servicesdi.ContentServiceFactory(contentRepository)
	classService := servicesdi.ClassServiceFactory(classrepository)

	return application.NewApplication(
		userService,
		adminService,
		teacherService,
		studentService,
		guardianService,
		contentService,
		classService,
	)
}

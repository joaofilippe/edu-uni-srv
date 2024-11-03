package application

import "github.com/joaofilippe/edu-uni-srv/core/services"

type Application struct {
	userService     services.IUserService
	adminService    services.IAdminService
	teacherService  services.ITeacherService
	studentService  services.IStudentService
	guardianService services.IGuardianService
}

func NewApplication(
	userService services.IUserService,
	adminService services.IAdminService,
	teacherService services.ITeacherService,
	studentService services.IStudentService,
	guardianService services.IGuardianService,
) *Application {
	return &Application{
		userService:     userService,
		adminService:    adminService,
		teacherService:  teacherService,
		studentService:  studentService,
		guardianService: guardianService,
	}
}

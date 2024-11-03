package application

import "github.com/joaofilippe/edu-uni-srv/core/services"

type Application struct {
	userService     iservices.IUserService
	adminService    iservices.IAdminService
	teacherService  iservices.ITeacherService
	studentService  iservices.IStudentService
	guardianService iservices.IGuardianService
}

func NewApplication(
	userService iservices.IUserService,
	adminService iservices.IAdminService,
	teacherService iservices.ITeacherService,
	studentService iservices.IStudentService,
	guardianService iservices.IGuardianService,
) *Application {
	return &Application{
		userService:     userService,
		adminService:    adminService,
		teacherService:  teacherService,
		studentService:  studentService,
		guardianService: guardianService,
	}
}

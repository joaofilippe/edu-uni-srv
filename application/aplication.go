package application

import "github.com/joaofilippe/edu-uni-srv/core/services"

type Application struct {
	UserService services.IUserService
	AdminService services.IAdminService
	TeacherService services.ITeacherService
	StudentService services.IStudentService
	GuardianService services.IGuardianService
}
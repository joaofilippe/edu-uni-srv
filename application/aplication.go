package application

import (
	"github.com/joaofilippe/edu-uni-srv/domain/services"
)

type Application struct {
	userService     iservices.IUserService
	adminService    iservices.IAdminService
	teacherService  iservices.ITeacherService
	studentService  iservices.IStudentService
	guardianService iservices.IGuardianService
	contentService iservices.IContentService
}

func NewApplication(
	userService iservices.IUserService,
	adminService iservices.IAdminService,
	teacherService iservices.ITeacherService,
	studentService iservices.IStudentService,
	guardianService iservices.IGuardianService,
	contentService iservices.IContentService,
) *Application {
	return &Application{
		userService:     userService,
		adminService:    adminService,
		teacherService:  teacherService,
		studentService:  studentService,
		guardianService: guardianService,
		contentService: contentService,
	}
}

func (a *Application) UserService() iservices.IUserService {
	return a.userService
}

func (a *Application) SetUserService(userService iservices.IUserService) {
	a.userService = userService
}

func (a *Application) AdminService() iservices.IAdminService {
	return a.adminService
}

func (a *Application) TeacherService() iservices.ITeacherService {
	return a.teacherService
}

func (a *Application) StudentService() iservices.IStudentService {
	return a.studentService
}

func (a *Application) GuardianService() iservices.IGuardianService {
	return a.guardianService
}

func (a *Application) ContentService() iservices.IContentService {
	return a.contentService
}

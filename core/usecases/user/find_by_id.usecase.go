package user

import (
	"github.com/google/uuid"

	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	studentEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	teacherEntities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/interfaces"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	UserRepo     repositories.IUserRepo
	AdminRepo    repositories.IAdminRepo
	GuardianRepo repositories.IGuardianRepo
	TeacherRepo  repositories.ITeacherRepo
	StudentRepo  repositories.IStudentRepo
}

func NewFindByIDUseCase(
	userRepo *repositories.IUserRepo,
	adminRepo *repositories.IAdminRepo,
	guardianRepo *repositories.IGuardianRepo,
	teacherRepo *repositories.ITeacherRepo,
	studentRepo *repositories.IStudentRepo,
) *FindByIDUseCase {
	return &FindByIDUseCase{
		UserRepo:     *userRepo,
		AdminRepo:    *adminRepo,
		GuardianRepo: *guardianRepo,
		TeacherRepo:  *teacherRepo,
		StudentRepo:  *studentRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*userEntities.User, error) {
	if id == uuid.Nil {
		return nil, commonErrors.ErrUserIDNot
	}

	user, err := u.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	userDetails, err := u.findUserDetails(*user)
	if err != nil {
		return nil, err
	}

	user.SetUserDetails(userDetails)

	return user, nil
}

func (u *FindByIDUseCase) findAdmin(id uuid.UUID) (*adminEntities.Admin, error) {
	admin, err := u.AdminRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (u *FindByIDUseCase) findGuardian(id uuid.UUID) (*guardianEntities.Guardian, error) {
	guardian, err := u.GuardianRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return guardian, nil
}

func (u *FindByIDUseCase) findTeacher(id uuid.UUID) (*teacherEntities.Teacher, error) {
	teacher, err := u.TeacherRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return teacher, nil
}

func (u *FindByIDUseCase) findStudent(id uuid.UUID) (*studentEntities.Student, error) {
	student, err := u.StudentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (u *FindByIDUseCase) findUserDetails(user userEntities.User) (interfaces.IUserDetails, error) {
	switch user.UserType() {
	case enums.Administrator:
		return u.findAdmin(user.ID())
	case enums.Guardian:
		return u.findGuardian(user.ID())
	case enums.Teacher:
		return u.findTeacher(user.ID())
	case enums.Student:
		return u.findStudent(user.ID())
	default:
		return nil, commonErrors.ErrUserTypeNotFound
	}
}

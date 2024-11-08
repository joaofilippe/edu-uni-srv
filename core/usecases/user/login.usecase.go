package userusecases

import (
	"github.com/google/uuid"

	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	"github.com/joaofilippe/edu-uni-srv/common/password"
	tokengenerator "github.com/joaofilippe/edu-uni-srv/common/token"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/interfaces"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type LoginUseCase struct {
	userRepository     irepositories.IUserRepo
	adminRepository    irepositories.IAdminRepo
	studentRepository  irepositories.IStudentRepo
	teacherRepository  irepositories.ITeacherRepo
	guardianRepository irepositories.IGuardianRepo
}

func NewLoginUseCase(
	userRepository *irepositories.IUserRepo,
	adminRepository *irepositories.IAdminRepo,
	studentRepository *irepositories.IStudentRepo,
	teacherRepository *irepositories.ITeacherRepo,
	guardianRepository *irepositories.IGuardianRepo,
) *LoginUseCase {
	return &LoginUseCase{
		*userRepository,
		*adminRepository,
		*studentRepository,
		*teacherRepository,
		*guardianRepository,
	}
}

func (u *LoginUseCase) Execute(email, givenPassword string) (string, error) {
	if email == "" || givenPassword == "" {
		return "", usecaseerrors.ErrUserInvalidEmail
	}

	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	if !password.CheckPasswordHash(givenPassword, user.Password()) {
		return "", usecaseerrors.ErrUserInvalidPassword
	}

	userDetails, err := u.getUserDetails(user.UserType(), user.ID())
	if err != nil {
		return "", err
	}

	if userDetails == nil {
		return "", usecaseerrors.ErrUserDetailsNotFound
	}

	user.SetUserDetails(&userDetails)

	token, err := tokengenerator.GenerateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *LoginUseCase) getUserDetails(userType enums.UserType, userID uuid.UUID) (interfaces.IUserDetails, error) {
	switch userType {
	case enums.Administrator:
		return u.adminRepository.FindByUserID(userID)
	case enums.Student:
		return u.studentRepository.FindByUserID(userID)
	case enums.Teacher:
		return u.teacherRepository.FindByUserID(userID)
	case enums.Guardian:
		return u.guardianRepository.FindByUserID(userID)
	}

	return nil, nil
}

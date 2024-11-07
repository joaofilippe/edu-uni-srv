package userusecases

import (
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	"github.com/joaofilippe/edu-uni-srv/common/password"
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

func (u *LoginUseCase) Execute(email, givenPassword string) error {
	if email == "" || givenPassword == "" {
		return usecaseerrors.ErrUserInvalidEmail
	}

	user, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return err
	}

	if password.CheckPasswordHash(givenPassword, user.Password()) {
		return usecaseerrors.ErrUserInvalidPassword
	}

	return nil
}

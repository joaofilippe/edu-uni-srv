package userusecase

import (
	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByEmailUseCase struct {
	userRepository irepositories.IUserRepo
}

func NewFindByEmailUseCase(userRepository irepositories.IUserRepo) *FindByEmailUseCase {
	return &FindByEmailUseCase{userRepository}
}

func (u *FindByEmailUseCase) Execute(email string) (*userEntities.User, error) {
	if email == "" {
		return &userEntities.User{}, commonErrors.ErrUserInvalidEmail
	}

	users, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return &userEntities.User{}, err
	}

	return users, nil
}

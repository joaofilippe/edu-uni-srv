package userusecases

import (
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUseCase struct {
	userRepository irepositories.IUserRepo
}

func NewFindAllUseCase(userRepository *irepositories.IUserRepo) *FindAllUseCase {
	return &FindAllUseCase{ *userRepository}
}

func (u *FindAllUseCase) Execute() ([]*userEntities.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return []*userEntities.User{}, err
	}

	return users, nil
}

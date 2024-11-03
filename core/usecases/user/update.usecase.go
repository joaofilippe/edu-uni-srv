package userusecase

import (
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type UpdateUseCase struct {
	UserRepo irepositories.IUserRepo
}

func NewUpdateUseCase(userRepo *irepositories.IUserRepo) *UpdateUseCase {
	return &UpdateUseCase{
		UserRepo: *userRepo,
	}
}

func (u *UpdateUseCase) Execute(user *userEntities.User) error {
	if user == nil {
		return nil
	}

	return nil
}
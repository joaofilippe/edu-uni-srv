package user

import (
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type UpdateUseCase struct {
	UserRepo repositories.IUserRepo
}

func NewUpdateUseCase(userRepo *repositories.IUserRepo) *UpdateUseCase {
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
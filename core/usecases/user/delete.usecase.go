package userusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type DeleteUseCase struct {
	UserRepo *irepositories.IUserRepo
}

func NewDeleteUseCase(userRepo *irepositories.IUserRepo) *DeleteUseCase {
	return &DeleteUseCase{
		UserRepo: userRepo,
	}
}

func (u *DeleteUseCase) Execute(id uuid.UUID) error {
	if id == uuid.Nil {
		return nil
	}

	return nil
}
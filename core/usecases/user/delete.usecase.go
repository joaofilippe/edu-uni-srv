package userusecase

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type DeleteUseCase struct {
	UserRepo *repositories.IUserRepo
}

func NewDeleteUseCase(userRepo *repositories.IUserRepo) *DeleteUseCase {
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
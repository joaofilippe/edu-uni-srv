package teacherusecase

import (
	"github.com/google/uuid"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type DeleteUseCase struct {
	deleteRepo irepositories.ITeacherRepo
}

func NewDeleteUseCase(deleteRepo irepositories.ITeacherRepo) *DeleteUseCase {
	return &DeleteUseCase{
		deleteRepo,
	}
}

func (uc *DeleteUseCase) Execute(id uuid.UUID) error {
	return uc.deleteRepo.Delete(id)
}

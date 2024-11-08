package guardianusecases

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type DeleteUseCase struct {
	guardianRepo irepositories.IGuardianRepo
}

func NewDeleteUseCase(guardianRepo *irepositories.IGuardianRepo) *DeleteUseCase {
	return &DeleteUseCase{
		guardianRepo: *guardianRepo,
	}
}

func (u *DeleteUseCase) Execute(id uuid.UUID) error {
	return u.guardianRepo.Delete(id)
}

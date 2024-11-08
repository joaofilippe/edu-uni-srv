package adminusecases

import (
	"github.com/google/uuid"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type UpdateUseCase struct {
	iAdminRepo irepositories.IAdminRepo
}

func NewUpdateUseCase(iAdminRepo *irepositories.IAdminRepo) *UpdateUseCase {
	return &UpdateUseCase{
		iAdminRepo: *iAdminRepo,
	}
}

func (uc *UpdateUseCase) Execute(id uuid.UUID) error {
	return uc.iAdminRepo.Update(id)
}

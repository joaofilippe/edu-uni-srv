package adminusecases

import (
	"github.com/google/uuid"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type DeleteUseCase struct {
	adminRepo irepositories.IAdminRepo
}

func NewDeleteUseCase(adminRepo irepositories.IAdminRepo) *DeleteUseCase {
	return &DeleteUseCase{
		adminRepo: adminRepo,
	}
}

func (uc *DeleteUseCase) Execute(id uuid.UUID) error {
	return uc.adminRepo.Delete(id)
}
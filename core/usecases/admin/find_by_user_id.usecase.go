package adminusecases

import (
	"github.com/google/uuid"

	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByUserIDUseCase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindByUserIDUseCase(adminRepo irepositories.IAdminRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		AdminRepo: adminRepo,
	}
}

func (u *FindByUserIDUseCase) Execute(id uuid.UUID) (*adminEntities.Admin, error) {
	if id == uuid.Nil {
		return nil, commonErrors.ErrAdminIDNot
	}

	return u.AdminRepo.FindByID(id)
}

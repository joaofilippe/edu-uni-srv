package guardian

import (
	"github.com/google/uuid"

	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	AdminRepo repositories.IAdminRepo
}

func NewFindByIDUseCase(adminRepo repositories.IAdminRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		AdminRepo: adminRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*adminEntities.Admin, error) {
	if id == uuid.Nil {
		return nil, commonErrors.ErrAdminIDNot
	}

	return u.AdminRepo.FindByID(id)
}

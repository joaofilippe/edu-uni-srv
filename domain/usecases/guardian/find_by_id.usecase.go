package guardianusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	guardianentities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	guardianRepo irepositories.IGuardianRepo
}

func NewFindByIDUseCase(guardianRepo *irepositories.IGuardianRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		guardianRepo: *guardianRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*guardianentities.Guardian, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.guardianRepo.FindByID(id)
}

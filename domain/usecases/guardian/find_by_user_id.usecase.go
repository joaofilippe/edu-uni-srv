package guardianusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	guardianentities "github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type FindByUserIDUseCase struct {
	guardianRepo irepositories.IGuardianRepo
}

func NewFindByUserIDUseCase(guardianRepo *irepositories.IGuardianRepo) *FindByUserIDUseCase {
	return &FindByUserIDUseCase{
		guardianRepo: *guardianRepo,
	}
}

func (u *FindByUserIDUseCase) Execute(id uuid.UUID) (*guardianentities.Guardian, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.guardianRepo.FindByID(id)
}

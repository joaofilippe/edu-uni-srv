package adminusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByUserIDUseCase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindByUserIDUseCase(adminRepo *irepositories.IAdminRepo) *FindByUserIDUseCase {
	return &FindByUserIDUseCase{
		AdminRepo: *adminRepo,
	}
}

func (u *FindByUserIDUseCase) Execute(id uuid.UUID) (*adminentities.Admin, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.AdminRepo.FindByID(id)
}

package adminusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindByIDUseCase(adminRepo *irepositories.IAdminRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		AdminRepo: *adminRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*adminentities.Admin, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.AdminRepo.FindByID(id)
}

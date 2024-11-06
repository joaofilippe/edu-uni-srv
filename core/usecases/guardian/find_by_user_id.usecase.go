package guardianusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common"
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByUserIDUseCase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindByUserIDUseCase(adminRepo irepositories.IAdminRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		AdminRepo: adminRepo,
	}
}

func (u *FindByUserIDUseCase) Execute(id uuid.UUID) (*adminentities.Admin, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.AdminRepo.FindByID(id)
}

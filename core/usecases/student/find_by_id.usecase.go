package studentusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common"
	studententities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	AdminRepo irepositories.IStudentRepo
}

func NewFindByIDUseCase(adminRepo *irepositories.IStudentRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		AdminRepo: *adminRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*studententities.Student, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot

	}

	return u.AdminRepo.FindByID(id)
}

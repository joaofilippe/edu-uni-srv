package studentusecase

import (
	"github.com/google/uuid"

	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"

	"github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
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
		return nil, commonErrors.ErrAdminIDNot

	}

	return u.AdminRepo.FindByID(id)
}

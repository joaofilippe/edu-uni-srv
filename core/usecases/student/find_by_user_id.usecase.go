package studentusecases

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/common"
	"github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByUserIDUseCase struct {
	StudentRepo irepositories.IStudentRepo
}

func NewFindByUserIDUseCase(studentRepo *irepositories.IStudentRepo) *FindByUserIDUseCase {
	return &FindByUserIDUseCase{
		StudentRepo: *studentRepo,
	}
}

func (u *FindByUserIDUseCase) Execute(id uuid.UUID) (*studententities.Student, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.StudentRepo.FindByID(id)
}

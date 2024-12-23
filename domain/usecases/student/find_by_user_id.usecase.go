package studentusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/common/errors"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
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
		return nil, errors.ErrAdminIDNot
	}

	return u.StudentRepo.FindByID(id)
}

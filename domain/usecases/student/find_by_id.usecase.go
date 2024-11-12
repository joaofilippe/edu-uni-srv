package studentusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	studententities "github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type FindByIDUseCase struct {
	studentRepository irepositories.IStudentRepo
}

func NewFindByIDUseCase(studentRepository *irepositories.IStudentRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		studentRepository: *studentRepository,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*studententities.Student, error) {
	if id == uuid.Nil {
		return nil, usecaseerrors.ErrAdminIDNot

	}

	return u.studentRepository.FindByID(id)
}

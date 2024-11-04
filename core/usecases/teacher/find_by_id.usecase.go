package teacherusecase

import (
	"github.com/google/uuid"

	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	teacherentities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	teacherRepo irepositories.ITeacherRepo
}

func NewFindByIDUseCase(teacherRepo irepositories.ITeacherRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		teacherRepo: teacherRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*teacherentities.Teacher, error) {
	if id == uuid.Nil {
		return nil, commonErrors.ErrAdminIDNot
	}

	return u.teacherRepo.FindByID(id)
}

package teacherusecase

import (
	"github.com/google/uuid"

	commonErrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	teacherentities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByUserIDUseCase struct {
	teacherRepo irepositories.ITeacherRepo
}

func NewFindByUserIDUseCase(teacherRepo *irepositories.ITeacherRepo) *FindByUserIDUseCase {
	return &FindByUserIDUseCase{
		teacherRepo: *teacherRepo,
	}
}

func (u *FindByUserIDUseCase) Execute(id uuid.UUID) (*teacherentities.Teacher, error) {
	if id == uuid.Nil {
		return nil, commonErrors.ErrAdminIDNot
	}

	return u.teacherRepo.FindByID(id)
}

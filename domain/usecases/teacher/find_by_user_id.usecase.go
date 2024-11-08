package teacherusecases

import (
	"github.com/google/uuid"
	usecaseerrors "github.com/joaofilippe/edu-uni-srv/common/errors"
	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
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
		return nil, usecaseerrors.ErrAdminIDNot
	}

	return u.teacherRepo.FindByID(id)
}

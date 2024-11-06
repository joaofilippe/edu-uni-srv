package teacherusecases

import (
	"github.com/google/uuid"

	usecaseerros "github.com/joaofilippe/edu-uni-srv/common"
	teacherentities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindByIDUseCase struct {
	teacherRepo irepositories.ITeacherRepo
}

func NewFindByIDUseCase(teacherRepo *irepositories.ITeacherRepo) *FindByIDUseCase {
	return &FindByIDUseCase{
		teacherRepo: *teacherRepo,
	}
}

func (u *FindByIDUseCase) Execute(id uuid.UUID) (*teacherentities.Teacher, error) {
	if id == uuid.Nil {
		return nil, usecaseerros.ErrAdminIDNot
	}

	return u.teacherRepo.FindByID(id)
}

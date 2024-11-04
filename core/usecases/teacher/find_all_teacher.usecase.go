package teacherusecase

import (
	"github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUseCase struct {
	teacherRepo irepositories.ITeacherRepo
}

func NewFindAllUseCase(teacherRepo *irepositories.ITeacherRepo) *FindAllUseCase {
	return &FindAllUseCase{
		teacherRepo: *teacherRepo,
	}
}

func (u *FindAllUseCase) Execute() ([]*teacherentities.Teacher, error) {
	return u.teacherRepo.FindAll()
}

package teacherusecases

import (
	"github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
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

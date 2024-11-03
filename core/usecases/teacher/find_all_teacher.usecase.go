package teacherusecase

import (
	teacherEntities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllTeacherUsecase struct {
	TeacherRepo irepositories.ITeacherRepo
}

func NewFindAllTeacherUsecase(teacherRepo irepositories.ITeacherRepo) *FindAllTeacherUsecase {
	return &FindAllTeacherUsecase{
		TeacherRepo: teacherRepo,
	}
}

func (u *FindAllTeacherUsecase) Execute() ([]*teacherEntities.Teacher, error) {
	return u.TeacherRepo.FindAll()
}

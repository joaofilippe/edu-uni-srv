package teacherusecase

import (
	teacherEntities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllTeacherUsecase struct {
	TeacherRepo repositories.ITeacherRepo
}

func NewFindAllTeacherUsecase(teacherRepo repositories.ITeacherRepo) *FindAllTeacherUsecase {
	return &FindAllTeacherUsecase{
		TeacherRepo: teacherRepo,
	}
}

func (u *FindAllTeacherUsecase) Execute() ([]*teacherEntities.Teacher, error) {
	return u.TeacherRepo.FindAll()
}

package teacher

import (
	"github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateTeacherUseCase struct {
	teacherRepo repositories.ITeacherRepo
}

func NewCreateTeacherUseCase(teacherRepo repositories.ITeacherRepo) *CreateTeacherUseCase {
	return &CreateTeacherUseCase{
		teacherRepo: teacherRepo,
	}
}

func (uc *CreateTeacherUseCase) Execute(teacher *teacher.CreateTeacher) error {
	return uc.teacherRepo.Save(teacher)
}
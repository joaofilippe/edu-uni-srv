package teacherusecases

import (
	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type UpdateUseCase struct {
	teacherRepo irepositories.ITeacherRepo
}

func NewUpdateUseCase(teacherRepo *irepositories.ITeacherRepo) *UpdateUseCase {
	return &UpdateUseCase{
		*teacherRepo,
	}
}

func (uc *UpdateUseCase) Execute(teacher *teacherentities.Teacher) error {
	return uc.teacherRepo.Update(teacher)
}

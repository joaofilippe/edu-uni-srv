package studentusecase

import (
	studententities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUseCase struct {
	StudentRepo repositories.IStudentRepo
}

func NewFindAllUseCase(studentRepo *repositories.IStudentRepo) *FindAllUseCase {
	return &FindAllUseCase{
		StudentRepo: *studentRepo,
	}
}

func (u *FindAllUseCase) Execute() ([]*studententities.Student, error) {
	return u.StudentRepo.FindAll()
}

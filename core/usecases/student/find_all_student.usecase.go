package studentUseCases

import (
	studentEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUsecase struct {
	StudentRepo repositories.IStudentRepo
}

func NewFindAllUsecase(studentRepo repositories.IStudentRepo) *FindAllUsecase {
	return &FindAllUsecase{
		StudentRepo: studentRepo,
	}
}

func (u *FindAllUsecase) Execute() ([]*studentEntities.Student, error) {
	return u.StudentRepo.FindAll()
}

package student

import (
	studentEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllStudentUsecase struct {
	StudentRepo repositories.IStudentRepo
}

func NewFindAllStudentUsecase(studentRepo repositories.IStudentRepo) *FindAllStudentUsecase {
	return &FindAllStudentUsecase{
		StudentRepo: studentRepo,
	}
}

func (u *FindAllStudentUsecase) Execute() ([]*studentEntities.Student, error) {
	return u.StudentRepo.FindAll()
}
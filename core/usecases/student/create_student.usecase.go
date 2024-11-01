package student

import (
	"github.com/google/uuid"
	studentEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateStudentUseCase struct {
	studentRepository repositories.IStudentRepo
}

func NewCreateStudentUseCase(studentRepository repositories.IStudentRepo) *CreateStudentUseCase {
	return &CreateStudentUseCase{
		studentRepository: studentRepository,
	}
}

func (uc *CreateStudentUseCase) Execute(createStudent *studentEntities.CreateStudent) error {
	if createStudent.EmptyID() {
		createStudent.SetId(uuid.New())
	}

	err := uc.studentRepository.Save(createStudent)
	if err != nil {
		return err
	}

	return nil
}

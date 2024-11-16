package studentusecases

import (
	"github.com/google/uuid"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type EnrollUseCase struct {
	studentRepository irepositories.IStudentRepo
}

func NewEnrollUseCase(studentRepository *irepositories.IStudentRepo) *EnrollUseCase {
	return &EnrollUseCase{
		studentRepository: *studentRepository,
	}
}

func (e *EnrollUseCase) Execute(studentID, classID uuid.UUID) error {
	student, err := e.studentRepository.FindByID(studentID)
	if err != nil {
		return err
	}

	student.SetClassesIDs(append(student.ClassesIDs(), classID))

	return e.studentRepository.Update(student)
}

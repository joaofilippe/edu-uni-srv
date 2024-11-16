package studentusecases

import (
	"github.com/google/uuid"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type EnrollUsecase struct {
	studentRepository irepositories.IStudentRepo
}

func NewEnrollUsecase(studentRepository irepositories.IStudentRepo) *EnrollUsecase {
	return &EnrollUsecase{
		studentRepository: studentRepository,
	}
}

func (e *EnrollUsecase) Enroll(studentID, classID uuid.UUID) error {
	student, err := e.studentRepository.FindByID(studentID)
	if err != nil {
		return err
	}

	student.SetClassesIDs(append(student.ClassesIDs(), classID))

	return e.studentRepository.Update(student)
}

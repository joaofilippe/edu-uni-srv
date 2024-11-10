package studentusecases

import (
	"github.com/google/uuid"

	studententities "github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type CreateUseCase struct {
	studentRepository irepositories.IStudentRepo
	userRepository    irepositories.IUserRepo
}

func NewCreateUseCase(
	studentRepository *irepositories.IStudentRepo,
	userRepository *irepositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*studentRepository,
		*userRepository,
	}
}

func (uc *CreateUseCase) Execute(createStudent *studententities.CreateStudent) (uuid.UUID, error) {
	createStudent.SetId(uuid.New())

	err := uc.studentRepository.Save(createStudent)
	if err != nil {
		err = uc.userRepository.Delete(createStudent.UserID())
		if err != nil {
			return uuid.UUID{}, err
		}

		return uuid.UUID{}, err
	}

	return createStudent.ID(), nil
}

package studentusecases

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/core/enums"
	studententities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
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
	if createStudent.EmptyID() {
		createStudent.SetId(uuid.New())
	}

	err := uc.studentRepository.Save(createStudent)
	if err != nil {
		return uuid.UUID{}, err
	}

	user, err := uc.userRepository.FindByID(createStudent.UserID())

	studentType := enums.Student

	newUser := user.CopyWith(
		nil,
		nil,
		nil,
		&studentType,
		nil,
	)

	err = uc.userRepository.Update(newUser)
	if err != nil {

		return uuid.UUID{}, err
	}

	return createStudent.ID(), nil
}

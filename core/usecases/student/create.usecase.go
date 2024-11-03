package studentusecase

import (
	"github.com/google/uuid"

	studentEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	studentRepository repositories.IStudentRepo
	userRepository    repositories.IUserRepo
}

func NewCreateUseCase(
	studentRepository *repositories.IStudentRepo,
	userRepository *repositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*studentRepository,
		*userRepository,
	}
}

func (uc *CreateUseCase) Execute(createStudent *studentEntities.CreateStudent) (uuid.UUID, error) {
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

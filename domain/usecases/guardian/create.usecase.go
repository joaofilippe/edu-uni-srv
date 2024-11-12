package guardianusecases

import (
	"errors"

	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type CreateUseCase struct {
	guardianRepository irepositories.IGuardianRepo
	studentRepository  irepositories.IStudentRepo
	userRepository     irepositories.IUserRepo
}

func NewCreateUseCase(
	guardianRepository *irepositories.IGuardianRepo,
	studentRepository *irepositories.IStudentRepo,
	userRepository *irepositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*guardianRepository,
		*studentRepository,
		*userRepository,
	}
}

func (c *CreateUseCase) Execute(
	guardian *guardianentities.CreateGuardian,
) (uuid.UUID, error) {
	student, err := c.studentRepository.FindByID(guardian.StudentID())
	if err != nil {
		return uuid.UUID{}, err
	}

	if student == nil {
		return uuid.UUID{}, errors.New("student not found")
	}

	guardian.SetId(uuid.New())
	err = c.guardianRepository.Save(guardian)
	if err != nil {
		err = c.userRepository.Delete(guardian.UserID())
		if err != nil {
			return uuid.UUID{}, err
		}
		return uuid.UUID{}, err
	}

	return guardian.ID(), nil
}

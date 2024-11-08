package guardianusecases

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	guardianRepository irepositories.IGuardianRepo
	userRepository     irepositories.IUserRepo
}

func NewCreateUseCase(
	guardianRepository *irepositories.IGuardianRepo,
	userRepository *irepositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*guardianRepository,
		*userRepository,
	}
}

func (c *CreateUseCase) Execute(
	guardian *guardianentities.CreateGuardian,
) (uuid.UUID, error) {
	if guardian.EmptyID() {
		guardian.SetId(uuid.New())
	}

	err := c.guardianRepository.Save(guardian)
	if err != nil {
		return uuid.UUID{}, err
	}

	user, err := c.userRepository.FindByID(guardian.UserID())
	if err != nil {
		return uuid.UUID{}, err
	}

	guardianType := enums.Guardian

	newUser := user.CopyWith(
		nil,
		nil,
		nil,
		&guardianType,
		nil,
	)

	err = c.userRepository.Update(newUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return guardian.ID(), nil
}

package admin

import (
	"github.com/google/uuid"

	adminEntites "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	adminRepository repositories.IAdminRepo
	userRepository  repositories.IUserRepo
}

func NewCreateUseCase(
	adminRepository repositories.IAdminRepo,
	userRepository repositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		adminRepository,
		userRepository,
	}
}

func (uc *CreateUseCase) Execute(createAdmin *adminEntites.CreateAdmin) (uuid.UUID, error) {
	if createAdmin.EmptyID() {
		createAdmin.SetId(uuid.New())
	}

	err := uc.adminRepository.Save(createAdmin)
	if err != nil {
		return uuid.UUID{}, err
	}

	user, err := uc.userRepository.FindByID(createAdmin.UserID())
	if err != nil {
		return uuid.UUID{}, err
	}

	administratorType := enums.Administrator

	newUser := user.CopyWith(
		nil,
		nil,
		nil,
		&administratorType,
		nil,
	)

	err = uc.userRepository.Update(newUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return createAdmin.UserID(), nil
}
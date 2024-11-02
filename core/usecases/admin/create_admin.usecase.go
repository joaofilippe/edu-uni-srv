package admin

import (
	"github.com/google/uuid"

	adminEntites "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateAdminUseCase struct {
	adminRepository repositories.IAdminRepo
	userRepository  repositories.IUserRepo
}

func NewCreateAdminUseCase(
	adminRepository repositories.IAdminRepo,
	userRepository repositories.IUserRepo,
) *CreateAdminUseCase {
	return &CreateAdminUseCase{
		adminRepository,
		userRepository,
	}
}

func (uc *CreateAdminUseCase) Execute(createAdmin *adminEntites.CreateAdmin) (uuid.UUID, error) {
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

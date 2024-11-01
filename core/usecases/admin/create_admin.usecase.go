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
		id := uuid.New()
		createAdmin.SetId(id)
	}

	err := uc.adminRepository.Save(createAdmin)
	if err != nil {
		return uuid.UUID{}, err
	}

	user, err := uc.userRepository.FindByID(createAdmin.UserID())
	if err != nil {
		return uuid.UUID{}, err
	}

	user.CopyWith(
		nil,
		nil,
		nil,
		enums.Administrator,
		
	)

	return createAdmin.UserID(), nil
}

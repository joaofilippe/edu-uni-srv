package adminusecases

import (
	"github.com/google/uuid"

	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	adminRepository irepositories.IAdminRepo
	userRepository  irepositories.IUserRepo
}

func NewCreateUseCase(
	adminRepository *irepositories.IAdminRepo,
	userRepository *irepositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*adminRepository,
		*userRepository,
	}
}

func (uc *CreateUseCase) Execute(createAdmin *adminentities.Create) (uuid.UUID, error) {
	createAdmin.SetId(uuid.New())

	err := uc.adminRepository.Save(createAdmin)
	if err != nil {
		return uuid.UUID{}, err
	}

	return createAdmin.UserID(), nil
}

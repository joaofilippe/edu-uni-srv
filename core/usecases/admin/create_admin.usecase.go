package admin

import (
	"github.com/google/uuid"

	adminEntites "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateAdminUseCase struct {
	adminRepository repositories.IAdminRepo
}

func NewCreateAdminUseCase(adminRepository repositories.IAdminRepo) *CreateAdminUseCase {
	return &CreateAdminUseCase{
		adminRepository: adminRepository,
	}
}

func (uc *CreateAdminUseCase) Execute(createAdmin *adminEntites.CreateAdmin) (uuid.UUID, error) {
	id := uuid.New()
	admin := adminEntites.NewAdmin(
		id,
		createAdmin.UserId(),
		createAdmin.Username(),
		createAdmin.Email(),
	)

	err := uc.adminRepository.Save(admin)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

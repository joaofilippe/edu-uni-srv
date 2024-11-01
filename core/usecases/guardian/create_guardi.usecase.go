package guardian

import (
	"github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateGuardianUseCase struct {
	guardianRepository repositories.IGuardianRepo
}

func NewCreateGuardianUseCase(
	guardianRepository repositories.IGuardianRepo,
) *CreateGuardianUseCase {
	return &CreateGuardianUseCase{guardianRepository}
}

func (c *CreateGuardianUseCase) Execute(
	guardian *guardian.CreateGuardian,
) error {
	return c.guardianRepository.Save(guardian)
}

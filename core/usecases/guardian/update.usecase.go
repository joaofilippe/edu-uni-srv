package guardianusecases

import (
	guardianentities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type UpdateUseCase struct {
	guardianRepo irepositories.IGuardianRepo
}

func NewUpdateUseCase(guardianRepo *irepositories.IGuardianRepo) *UpdateUseCase {
	return &UpdateUseCase{
		guardianRepo: *guardianRepo,
	}
}

func (u *UpdateUseCase) Execute(guardian *guardianentities.Guardian) error {
	return u.guardianRepo.Update(guardian)
}

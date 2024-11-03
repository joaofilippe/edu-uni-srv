package guardianusecase

import (
	guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllGuardianUsecase struct {
	GuardianRepo irepositories.IGuardianRepo
}

func NewFindAllGuardianUsecase(guardianRepo irepositories.IGuardianRepo) *FindAllGuardianUsecase {
	return &FindAllGuardianUsecase{
		GuardianRepo: guardianRepo,
	}
}

func (u *FindAllGuardianUsecase) Execute() ([]*guardianEntities.Guardian, error) {
	return u.GuardianRepo.FindAll()
}

package guardianusecases

import (
	guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUsecase struct {
	GuardianRepo irepositories.IGuardianRepo
}

func NewFindAllUsecase(guardianRepo irepositories.IGuardianRepo) *FindAllUsecase {
	return &FindAllUsecase{
		GuardianRepo: guardianRepo,
	}
}

func (u *FindAllUsecase) Execute() ([]*guardianEntities.Guardian, error) {
	return u.GuardianRepo.FindAll()
}

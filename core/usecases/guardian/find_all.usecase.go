package guardianusecases

import (
	guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUseCase struct {
	GuardianRepo irepositories.IGuardianRepo
}

func NewFindAllUseCase(guardianRepo irepositories.IGuardianRepo) *FindAllUseCase {
	return &FindAllUseCase{
		GuardianRepo: guardianRepo,
	}
}

func (u *FindAllUseCase) Execute() ([]*guardianEntities.Guardian, error) {
	return u.GuardianRepo.FindAll()
}

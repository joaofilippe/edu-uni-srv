package guardianusecases

import (
	guardianEntities "github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type FindAllUseCase struct {
	guardianRepo irepositories.IGuardianRepo
}

func NewFindAllUseCase(guardianRepo *irepositories.IGuardianRepo) *FindAllUseCase {
	return &FindAllUseCase{
		guardianRepo: *guardianRepo,
	}
}

func (u *FindAllUseCase) Execute() ([]*guardianEntities.Guardian, error) {
	return u.guardianRepo.FindAll()
}

package adminusecases

import (
	adminEntites "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUseCase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindAllUseCase(adminRepo irepositories.IAdminRepo) *FindAllUseCase {
	return &FindAllUseCase{
		AdminRepo: adminRepo,
	}
}

func (u *FindAllUseCase) Execute() ([]*adminEntites.Admin, error) {
	return u.AdminRepo.FindAll()
}

package adminusecases

import (
	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllUseCase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindAllUseCase(adminRepo *irepositories.IAdminRepo) *FindAllUseCase {
	return &FindAllUseCase{
		AdminRepo: *adminRepo,
	}
}

func (u *FindAllUseCase) Execute() ([]*adminentities.Admin, error) {
	return u.AdminRepo.FindAll()
}

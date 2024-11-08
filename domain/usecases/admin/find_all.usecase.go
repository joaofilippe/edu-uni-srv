package adminusecases

import (
	adminentities "github.com/joaofilippe/edu-uni-srv/domain/entities/admin"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
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

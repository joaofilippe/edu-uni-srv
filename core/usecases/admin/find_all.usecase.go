package adminusecase

import (
	adminEntites "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllAdminUsecase struct {
	AdminRepo irepositories.IAdminRepo
}

func NewFindAllAdminUsecase(adminRepo irepositories.IAdminRepo) *FindAllAdminUsecase {
	return &FindAllAdminUsecase{
		AdminRepo: adminRepo,
	}
}

func (u *FindAllAdminUsecase) Execute() ([]*adminEntites.Admin, error) {
	return u.AdminRepo.FindAll()
}

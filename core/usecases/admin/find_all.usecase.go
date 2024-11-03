package adminusecase

import (
	adminEntites "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type FindAllAdminUsecase struct {
	AdminRepo repositories.IAdminRepo
}

func NewFindAllAdminUsecase(adminRepo repositories.IAdminRepo) *FindAllAdminUsecase {
	return &FindAllAdminUsecase{
		AdminRepo: adminRepo,
	}
}

func (u *FindAllAdminUsecase) Execute() ([]*adminEntites.Admin, error) {
	return u.AdminRepo.FindAll()
}

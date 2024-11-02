package repositories

import (
	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type IAdminRepo interface {
	Save(admin *adminEntities.CreateAdmin) error
	FindAll() ([]*adminEntities.Admin, error)
	FindByID(id string) (*adminEntities.Admin, error)
	FindByEmail(email string) (*adminEntities.Admin, error)
	FindByUsername(username string) (*adminEntities.Admin, error)
}

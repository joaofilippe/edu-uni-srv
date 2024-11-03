package services

import (
	"github.com/google/uuid"

	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type IAdminService interface {
	Create(admin *adminEntities.Admin) error
	FindAll() ([]adminEntities.Admin, error)
	FindById(id string) (*adminEntities.Admin, error)
	FindByUserID(userID uuid.UUID) (*adminEntities.Admin, error)
	Update(admin *adminEntities.Admin) error
	Delete(id string) error
}

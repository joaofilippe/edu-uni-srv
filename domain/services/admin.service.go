package iservices

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/entities/admin"
)

type IAdminService interface {
	Create(admin *adminentities.Create) (uuid.UUID, error)
	FindAll() ([]*adminentities.Admin, error)
	FindById(id uuid.UUID) (*adminentities.Admin, error)
	FindByUserID(userID uuid.UUID) (*adminentities.Admin, error)
	Update(admin *adminentities.Admin) error
	Delete(id uuid.UUID) error
}

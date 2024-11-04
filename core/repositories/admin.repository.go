package irepositories

import (
	"github.com/google/uuid"
	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type IAdminRepo interface {
	Save(admin *adminEntities.CreateUseCase) error
	FindAll() ([]*adminEntities.Admin, error)
	FindByUserID(userID uuid.UUID) (*adminEntities.Admin, error)
	FindByID(id uuid.UUID) (*adminEntities.Admin, error)
	Update(id uuid.UUID) error
	Delete(id uuid.UUID) error
}

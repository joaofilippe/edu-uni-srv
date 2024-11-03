package irepositories

import (
	"github.com/google/uuid"
	guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
)

type IGuardianRepo interface {
	Save(guardian *guardianEntities.CreateGuardian) error
	FindAll() ([]*guardianEntities.Guardian, error)
	FindByID(id uuid.UUID) (*guardianEntities.Guardian, error)
	FindByEmail(email string) (*guardianEntities.Guardian, error)
	FindByUsername(username string) (*guardianEntities.Guardian, error)
}

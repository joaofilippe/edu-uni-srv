package irepositories

import (
	"github.com/google/uuid"
	guardianEntities "github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
)

type IGuardianRepo interface {
	Save(guardian *guardianEntities.CreateGuardian) error
	FindAll() ([]*guardianEntities.Guardian, error)
	FindByID(id uuid.UUID) (*guardianEntities.Guardian, error)
	FindByUserID(userID uuid.UUID) (*guardianEntities.Guardian, error)
	FindByEmail(email string) (*guardianEntities.Guardian, error)
	FindByUsername(username string) (*guardianEntities.Guardian, error)
	Update(guardian *guardianEntities.Guardian) error
	Delete(id uuid.UUID) error
}

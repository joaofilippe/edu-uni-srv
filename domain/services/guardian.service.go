package iservices

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
)

type IGuardianService interface {
	Create(guardian *guardianentities.CreateGuardian) (uuid.UUID, error)
	FindAll() ([]*guardianentities.Guardian, error)
	FindById(id uuid.UUID) (*guardianentities.Guardian, error)
	FindByUserID(userID uuid.UUID) (*guardianentities.Guardian, error)
	Update(guardian *guardianentities.Guardian) error
	Delete(id uuid.UUID) error
}

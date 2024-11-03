package services

import (
	"github.com/google/uuid"

	guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
)

type IGuardianService interface {
	Create(guardian *guardianEntities.Guardian) error
	FindAll() ([]guardianEntities.Guardian, error)
	FindById(id string) (*guardianEntities.Guardian, error)
	FindByUserID(userID uuid.UUID) (*guardianEntities.Guardian, error)
	Update(guardian *guardianEntities.Guardian) error
	Delete(id string) error
}

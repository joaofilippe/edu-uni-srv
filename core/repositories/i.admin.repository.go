package irepositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type IAdminRepo interface {
	Save(admin *adminentities.CreateUseCase) error
	FindAll() ([]*adminentities.Admin, error)
	FindByUserID(userID uuid.UUID) (*adminentities.Admin, error)
	FindByID(id uuid.UUID) (*adminentities.Admin, error)
	Update(id uuid.UUID) error
	Delete(id uuid.UUID) error
}
package services

import (
	"github.com/google/uuid"

	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
)

type IUserService interface {
	Create(user *userEntities.User) error
	FindAll() ([]userEntities.User, error)
	FindById(id uuid.UUID) (*userEntities.User, error)
	FindByEmail(email string) (*userEntities.User, error)
	FindByUsername(username string) (*userEntities.User, error)
	Update(user *userEntities.User) error
	Delete(id string) error
}

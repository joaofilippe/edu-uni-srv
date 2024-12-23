package iservices

import (
	"github.com/google/uuid"

	userEntities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
)

type IUserService interface {
	Create(user *userEntities.CreateUser) (uuid.UUID, error)
	Login(email, password string) (string, error)
	FindAll() ([]*userEntities.User, error)
	FindById(id uuid.UUID) (*userEntities.User, error)
	FindByEmail(email string) (*userEntities.User, error)
	Update(user *userEntities.User) error
	Delete(id string) error
}

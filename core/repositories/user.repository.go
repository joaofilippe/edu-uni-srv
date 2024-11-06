package irepositories

import (
	"github.com/google/uuid"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
)

type IUserRepo interface {
	Save(user *userEntities.CreateUser) error
	Update(user *userEntities.User) error
	FindAll() ([]*userEntities.User, error)
	FindByID(id uuid.UUID) (*userEntities.User, error)
	FindByEmail(email string) (*userEntities.User, error)
	FindByUsername(username string) (*userEntities.User, error)
}

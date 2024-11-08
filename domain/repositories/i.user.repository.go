package irepositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/user"
)

type IUserRepo interface {
	Save(user *userentities.CreateUser) error
	Update(user *userentities.User) error
	FindAll() ([]*userentities.User, error)
	FindByID(id uuid.UUID) (*userentities.User, error)
	FindByEmail(email string) (*userentities.User, error)
	FindByUsername(username string) (*userentities.User, error)
	Delete(id uuid.UUID) error
}

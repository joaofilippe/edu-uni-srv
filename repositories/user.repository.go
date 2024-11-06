package repositories

import (
	"github.com/google/uuid"

	userentities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type UserDatabase struct {
	connection *database.DBConnection
}

func NewUserDatabase(connection *database.DBConnection) *UserDatabase {
	return &UserDatabase{
		connection,
	}
}

func (u *UserDatabase) Save(user *userentities.CreateUser) error {
	return nil
}
func (u *UserDatabase) Update(user *userentities.User) error {
	return nil
}

func (u *UserDatabase) FindAll() ([]*userentities.User, error) {
	return nil, nil
}

func (u *UserDatabase) FindByID(id uuid.UUID) (*userentities.User, error) {
	return nil, nil
}

func (u *UserDatabase) FindByEmail(email string) (*userentities.User, error) {
	return nil, nil
}

func (u *UserDatabase) FindByUsername(username string) (*userentities.User, error) {
	return nil, nil
}

func (u *UserDatabase) Delete(id uuid.UUID) error {
	return nil
}

package repositories

import (
	"github.com/google/uuid"

	userentities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type UserRepository struct {
	connection *database.DBConnection
}

func NewUserRepository(connection *database.DBConnection) *UserRepository {
	return &UserRepository{
		connection,
	}
}

func (u *UserRepository) Save(user *userentities.CreateUser) error {
	tx := u.connection.DBConnection.MustBegin()
	tx.Exec("INSERT INTO users (id, username, password, email, user_type) VALUES ($1, $2, $3, $4, $5)", user.ID(), user.Username(), user.Password(), user.Email(), user.UserType())

	return nil
}
func (u *UserRepository) Update(user *userentities.User) error {
	return nil
}

func (u *UserRepository) FindAll() ([]*userentities.User, error) {
	return nil, nil
}

func (u *UserRepository) FindByID(id uuid.UUID) (*userentities.User, error) {
	return nil, nil
}

func (u *UserRepository) FindByEmail(email string) (*userentities.User, error) {
	return nil, nil
}

func (u *UserRepository) FindByUsername(username string) (*userentities.User, error) {
	return nil, nil
}

func (u *UserRepository) Delete(id uuid.UUID) error {
	return nil
}

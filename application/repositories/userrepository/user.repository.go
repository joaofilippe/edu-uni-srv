package user

import (
	"fmt"
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
	res, err := tx.Exec(SaveUserQuery, user.ID(), user.Email(), user.Password(), user.Username(), user.UserType().String())
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(res.RowsAffected())

	return tx.Commit()
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

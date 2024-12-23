package userrepository

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	userentities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
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
	userDB := fromCreateEntity(user)
	res, err := tx.NamedExec(SaveUserQuery, userDB)
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
	userDB := &UserDBModel{}
	err := u.connection.DBConnection.Get(userDB, FindByIDQuery, id)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return userDB.toEntity()
}

func (u *UserRepository) FindByEmail(email string) (*userentities.User, error) {
	userDB := &UserDBModel{}
	err := u.connection.DBConnection.Get(userDB, FindByEmailQuery, email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	return userDB.toEntity()
}

func (u *UserRepository) FindByUsername(username string) (*userentities.User, error) {
	return nil, nil
}

func (u *UserRepository) Delete(id uuid.UUID) error {
	tx := u.connection.DBConnection.MustBegin()

	_, err := tx.Exec(DeleteQuery, id)
	if err != nil {
		return tx.Rollback()
	}

	return tx.Commit()
}

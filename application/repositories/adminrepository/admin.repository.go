package adminrepository

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type AdminRepository struct {
	conn *database.DBConnection
}

func NewAdminRepository(conn *database.DBConnection) *AdminRepository {
	return &AdminRepository{conn}
}

func (a *AdminRepository) Save(admin *adminentities.Create) error {
	tx := a.conn.DBConnection.MustBegin()
	res, err := tx.Exec(SaveAdminQuery, admin.ID(), admin.UserID())
	if err != nil {
		return tx.Rollback()
	}

	fmt.Println(res)
	return tx.Commit()
}

func (a *AdminRepository) FindAll() ([]*adminentities.Admin, error) {
	return nil, nil
}

func (a *AdminRepository) FindByUserID(userID uuid.UUID) (*adminentities.Admin, error) {
	return nil, nil
}

func (a *AdminRepository) FindByID(id uuid.UUID) (*adminentities.Admin, error) {
	return nil, nil
}

func (a *AdminRepository) Update(id uuid.UUID) error {
	return nil
}

func (a *AdminRepository) Delete(id uuid.UUID) error {
	return nil
}
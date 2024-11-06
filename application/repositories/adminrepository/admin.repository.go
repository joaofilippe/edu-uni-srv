package repositories

import (
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
	return nil
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

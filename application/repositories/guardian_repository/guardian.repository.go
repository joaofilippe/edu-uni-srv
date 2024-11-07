package guardianrepository

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type GuardianRepository struct {
	conn *database.DBConnection
}

func NewGuardianRepository(conn *database.DBConnection) *GuardianRepository {
	return &GuardianRepository{conn}
}

func (g *GuardianRepository) Save(guardian *guardianentities.CreateGuardian) error {
	return nil
}

func (g *GuardianRepository) FindAll() ([]*guardianentities.Guardian, error) {
	return nil, nil
}

func (g *GuardianRepository) FindByID(id uuid.UUID) (*guardianentities.Guardian, error) {
	return nil, nil
}

func (g *GuardianRepository) FindByUserID(userID uuid.UUID) (*guardianentities.Guardian, error) {
	return nil, nil
}

func (g *GuardianRepository) FindByEmail(email string) (*guardianentities.Guardian, error) {
	return nil, nil
}

func (g *GuardianRepository) FindByUsername(username string) (*guardianentities.Guardian, error) {
	return nil, nil
}

func (g *GuardianRepository) Update(guardian *guardianentities.Guardian) error {
	return nil
}

func (g *GuardianRepository) Delete(id uuid.UUID) error {
	return nil
}

package classrepository

import (
	"github.com/joaofilippe/edu-uni-srv/domain/entities/class"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type ClassRepository struct {
	conn *database.DBConnection
}

func NewClassRepository(conn *database.DBConnection) *ClassRepository {
	return &ClassRepository{conn: conn}
}

func (c *ClassRepository) Save(class *classentities.CreateClass) error {
	tx := c.conn.DBConnection.MustBegin()
	createClassDB := &CreateClassDbModel{}
	createClassDB.fromEntity(class)

	_, err := tx.NamedExec(
		SaveQuery,
		createClassDB,
	)
	if err != nil {
		tx.Rollback()
		return err
	}
	
	return tx.Commit()
}
package contentrepository

import (
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	"github.com/joaofilippe/edu-uni-srv/infra/database"
)

type ContentRepository struct {
	conn *database.DBConnection
}

func NewContentRepository(conn *database.DBConnection) *ContentRepository {
	return &ContentRepository{conn: conn}
}

func (c *ContentRepository) Save(content *contententities.CreateContent) error {
	tx := c.conn.DBConnection.MustBegin()
	createContentDB := &CreateContentDbModel{}
	createContentDB.fromEntity(content)

	_, err := tx.NamedExec(
		SaveQuery,
		createContentDB,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}


func (c *ContentRepository) FindAll() ([]*contententities.Content, error) {
	var contentsDB []*ContentDbModel
	err := c.conn.DBConnection.Select(&contentsDB, FindAllQuery)
	if err != nil {
		return []*contententities.Content{}, err
	}

	var contents []*contententities.Content
	for _, content := range contentsDB {
		contents = append(contents, content.toEntity())
	}

	return contents, nil
}

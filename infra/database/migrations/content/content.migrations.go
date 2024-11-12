package contentmigrations

import "github.com/joaofilippe/edu-uni-srv/infra/database"

func CreateContentTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(upQuery)
	if err != nil {
		return err
	}

	return nil
}

func DropContentTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(downQuery)
	if err != nil {
		return err
	}

	return nil
}

package guardianmigrations

import "github.com/joaofilippe/edu-uni-srv/infra/database"

func CreateGuardiansTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(upQuery)
	if err != nil {
		return err
	}

	return nil
}

func DropGuardiansTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(downQuery)
	if err != nil {
		return err
	}

	return nil
}

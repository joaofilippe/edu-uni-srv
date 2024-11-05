package usermigrations

import "github.com/joaofilippe/edu-uni-srv/infra/database"

func CreateUsersTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(upQuery)
	if err != nil {
		return err
	}

	return nil
}

func DropUsersTable(conn *database.DBConnection) error {
	_, err := conn.DBConnection.Exec(downQuery)
	if err != nil {
		return err
	}

	return nil
}

package migrations

import (
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	usermigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/user"
)

func RunMigrations(connection *database.DBConnection) error {
	err := usermigrations.CreateUsersTable(connection)
	if err != nil {
		return err
	}

	err = usermigrations.DropUsersTable(connection)
	if err != nil {
		return err
	}

	return nil
}

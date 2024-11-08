package migrations

import (
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	adminmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/admin"
	guardianmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/guardian"
	studentmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/student"
	teachermigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/teacher"
	usermigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/user"
)

func RunMigrations(connection *database.DBConnection) error {
	err := usermigrations.CreateUsersTable(connection)
	if err != nil {
		return err
	}

	err = adminmigrations.CreateAdminsTable(connection)
	if err != nil {
		return err
	}

	err = studentmigrations.CreateStudentsTable(connection)
	if err != nil {
		return err
	}

	err = guardianmigrations.CreateGuardiansTable(connection)
	if err != nil {
		return err
	}

	err = teachermigrations.CreateTeachersTable(connection)
	if err != nil {
		return err
	}

	return nil
}

func createDatabase(connection *database.DBConnection){
	connection.DBConnection.Exec(`
	CREATE SERVER IF NOT EXISTS postgres`)
}
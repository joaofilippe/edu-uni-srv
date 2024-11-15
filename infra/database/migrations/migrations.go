package migrations

import (
	"github.com/joaofilippe/edu-uni-srv/infra/database"
	adminmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/admin"
	classmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/class"
	contentmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/content"
	guardianmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/guardian"
	psqlextensions "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/psql_extensions"
	studentmigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/student"
	teachermigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/teacher"
	usermigrations "github.com/joaofilippe/edu-uni-srv/infra/database/migrations/user"
)

func RunMigrations(connection *database.DBConnection) error {
	err := psqlextensions.CreatePsqlExtensionsQuery(connection)
	if err != nil {
		return err
	}

	err = usermigrations.CreateUsersTable(connection)
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

	err = contentmigrations.CreateContentTable(connection)
	if err != nil {
		return err
	}

	err = classmigrations.CreateClassTable(connection)
	if err != nil {
		return err
	}

	return err
}

func createDatabase(connection *database.DBConnection) {
	connection.DBConnection.Exec(`
	CREATE SERVER IF NOT EXISTS postgres`)
}

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

	err = usermigrations.DropUsersTable(connection)
	if err != nil {
		return err
	}

	err = adminmigrations.DropAdminsTable(connection)
	if err != nil {
		return err
	}

	err = studentmigrations.DropStudentsTable(connection)
	if err != nil {
		return err

	}

	err = guardianmigrations.DropGuardiansTable(connection)
	if err != nil {
		return err
	}

	err = teachermigrations.DropTeachersTable(connection)
	if err != nil {
		return err
	}

	return nil
}

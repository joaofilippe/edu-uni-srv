package repositories

import (
	"github.com/google/uuid"
	studentsEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
)

type IStudentRepo interface {
	Save(user *studentsEntities.CreateStudent) error
	FindAll() ([]*studentsEntities.Student, error)
	FindByID(id uuid.UUID) (*studentsEntities.Student, error)
	FindByEmail(email string) (*studentsEntities.Student, error)
	FindByUsername(username string) (*studentsEntities.Student, error)
}

package repositories

import (
	studentsEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
)

type IStudentRepo interface {
	Save(user *studentsEntities.CreateStudent) error
}

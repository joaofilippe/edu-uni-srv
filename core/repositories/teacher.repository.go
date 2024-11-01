package repositories

import "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"

type ITeacherRepo interface {
	Save(teacher *teacher.CreateTeacher) error
}

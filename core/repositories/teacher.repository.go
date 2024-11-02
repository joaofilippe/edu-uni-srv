package repositories

import "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"

type ITeacherRepo interface {
	Save(teacher *teacher.CreateTeacher) error
	FindAll() ([]*teacher.Teacher, error)
	FindByID(id string) (*teacher.Teacher, error)
	FindByEmail(email string) (*teacher.Teacher, error)
	FindByUsername(username string) (*teacher.Teacher, error)
}

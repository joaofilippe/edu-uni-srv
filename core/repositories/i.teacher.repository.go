package irepositories

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
)

type ITeacherRepo interface {
	Save(teacher *teacherentities.CreateTeacher) error
	FindAll() ([]*teacherentities.Teacher, error)
	FindByID(id uuid.UUID) (*teacherentities.Teacher, error)
	FindByEmail(email string) (*teacherentities.Teacher, error)
	FindByUsername(username string) (*teacherentities.Teacher, error)
	Update(teacher *teacherentities.Teacher) error
	Delete(id uuid.UUID) error
}

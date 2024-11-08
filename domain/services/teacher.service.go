package iservices

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
)

type ITeacherService interface {
	Create(teacher *teacherentities.Teacher) error
	FindAll() ([]*teacherentities.Teacher, error)
	FindById(id uuid.UUID) (*teacherentities.Teacher, error)
	FindByUserID(userID uuid.UUID) (*teacherentities.Teacher, error)
	Update(teacher *teacherentities.Teacher) error
	Delete(id uuid.UUID) error
}

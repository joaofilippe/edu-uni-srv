package services

import (
	"github.com/google/uuid"

	teacherEntities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
)

type ITeacherService interface {
	Create(teacher *teacherEntities.Teacher) error
	FindAll() ([]teacherEntities.Teacher, error)
	FindById(id string) (*teacherEntities.Teacher, error)
	FindByUserID(userID uuid.UUID) (*teacherEntities.Teacher, error)
	Update(teacher *teacherEntities.Teacher) error
	Delete(id string) error
}

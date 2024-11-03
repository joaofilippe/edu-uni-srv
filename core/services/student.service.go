package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/entities/student"
)

type IStudentService interface {
	Create(student *studententities.Student) error
	FindAll() ([]*studententities.Student, error)
	FindById(id uuid.UUID) (*studententities.Student, error)
	FindByUserID(userID uuid.UUID) (*studententities.Student, error)
	Update(student *studententities.Student) error
	Delete(id string) error
}

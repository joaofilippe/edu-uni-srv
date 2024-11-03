package services

import (
	"github.com/google/uuid"
	studentEntities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
)

type IStudentService interface {
	Create(student *studentEntities.Student) error
	FindAll() ([]studentEntities.Student, error)
	FindById(id string) (*studentEntities.Student, error)
	FindByUserID(userID uuid.UUID) (*studentEntities.Student, error)
	Update(student *studentEntities.Student) error
	Delete(id string) error
}
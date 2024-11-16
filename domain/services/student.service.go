package iservices

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/student"
)

type IStudentService interface {
	Create(student *studententities.CreateStudent) (uuid.UUID,error)
	Enroll(studentID, classID uuid.UUID) error
	FindAll() ([]*studententities.Student, error)
	FindById(id uuid.UUID) (*studententities.Student, error)
	FindByUserID(userID uuid.UUID) (*studententities.Student, error)
	Update(student *studententities.Student) error
	Delete(id string) error
}

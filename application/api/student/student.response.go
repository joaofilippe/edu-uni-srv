package studentweb

import (
	"time"

	"github.com/google/uuid"
	studententities "github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type StudentReponse struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Age          int
	ClassesIDs   []uuid.UUID
	GradesIDs    []uuid.UUID
	Disabilities []enums.Disability
	GuardianID   uuid.UUID
	Address      string
	Phone        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Active       bool
}

func (s *StudentReponse) fromEntity(entity *studententities.Student) {
	s.ID = entity.ID()
	s.UserID = entity.UserID()
	s.Age = entity.Age()
	s.ClassesIDs = entity.ClassesIDs()
	s.GradesIDs = entity.GradesIDs()
	s.Disabilities = entity.Disabilities()
	s.GuardianID = entity.GuardianID()
	s.Address = entity.Address()
	s.Phone = entity.Phone()
	s.CreatedAt = entity.CreatedAt()
	s.UpdatedAt = entity.UpdatedAt()
	s.Active = entity.Active()
}

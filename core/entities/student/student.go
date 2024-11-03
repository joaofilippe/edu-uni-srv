package studententities

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/core/entities/class"
	"github.com/joaofilippe/edu-uni-srv/core/entities/grade"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
)

type  Student struct {
	id           uuid.UUID
	userID       uuid.UUID
	name         string
	age          int
	classes      []class.Class
	grades       []grade.Grade
	disabilities []enums.Disability
	guardianID   uuid.UUID
	address      string
	phone        string
}

func NewStudent(
	id uuid.UUID,
	userID uuid.UUID,
	name string,
	age int,
	classes []class.Class,
	grades []grade.Grade,
	disabilities []enums.Disability,
	guardianID uuid.UUID,
	address string,
	phone string,
) *Student {
	return &Student{
		id,
		userID,
		name,
		age,
		classes,
		grades,
		disabilities,
		guardianID,
		address,
		phone,
	}
}

func (s *Student) Id() uuid.UUID {
	return s.id
}

func (s *Student) UserID() uuid.UUID {
	return s.userID
}

func (s *Student) Name() string {
	return s.name
}

func (s *Student) Age() int {
	return s.age
}

func (s *Student) Classes() []class.Class {
	return s.classes
}

func (s *Student) Grades() []grade.Grade {
	return s.grades
}

func (s *Student) Disabilities() []enums.Disability {
	return s.disabilities
}

func (s *Student) GuardianID() uuid.UUID {
	return s.guardianID
}

func (s *Student) Address() string {
	return s.address
}

func (s *Student) Phone() string {
	return s.phone
}

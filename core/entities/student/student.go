package studententities

import (
	"time"

	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/core/entities/class"
	"github.com/joaofilippe/edu-uni-srv/core/entities/grade"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
)

type Student struct {
	id           uuid.UUID
	userID       uuid.UUID
	age          int
	classes      []classentities.Class
	grades       []gradeentities.Grade
	disabilities []enums.Disability
	guardianID   uuid.UUID
	address      string
	phone        string
	createdAt    time.Time
	updatedAt    time.Time
	active       bool
}

func NewStudent(
	id uuid.UUID,
	userID uuid.UUID,
	age int,
	classes []classentities.Class,
	grades []gradeentities.Grade,
	disabilities []enums.Disability,
	guardianID uuid.UUID,
	address string,
	phone string,
	createdAt time.Time,
	updatedAt time.Time,
	active bool,
) *Student {
	return &Student{
		id,
		userID,
		age,
		classes,
		grades,
		disabilities,
		guardianID,
		address,
		phone,
		createdAt,
		updatedAt,
		active,
	}
}

func (s *Student) ID() uuid.UUID {
	return s.id
}

func (s *Student) UserID() uuid.UUID {
	return s.userID
}

func (s *Student) Age() int {
	return s.age
}

func (s *Student) Classes() []classentities.Class {
	return s.classes
}

func (s *Student) Grades() []gradeentities.Grade {
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

func (s *Student) CreatedAt() time.Time {
	return s.createdAt
}

func (s *Student) UpdatedAt() time.Time {
	return s.updatedAt
}

func (s *Student) Active() bool {
	return s.active
}

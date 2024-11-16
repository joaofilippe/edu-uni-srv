package studententities

import (
	"time"

	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/domain/enums"
)

type Student struct {
	id           uuid.UUID
	userID       uuid.UUID
	age          int
	classesIDs   []uuid.UUID
	gradesIDs    []uuid.UUID
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
	classesIDs []uuid.UUID,
	gradesIDs []uuid.UUID,
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
		classesIDs,
		gradesIDs,
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

func (s *Student) ClassesIDs() []uuid.UUID {
	return s.classesIDs
}

func (s *Student) GradesIDs() []uuid.UUID {
	return s.gradesIDs
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

func (s *Student) SetClassesIDs(classesIDs []uuid.UUID) {
	s.classesIDs = classesIDs
}

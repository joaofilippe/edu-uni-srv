package classentities

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	id         uuid.UUID
	name       string
	teacherID  uuid.UUID
	studentsID []uuid.UUID
	contentsID []uuid.UUID
	createdAt  time.Time
	updatedAt  time.Time
}

func NewClass(
	id uuid.UUID,
	name string,
	teacherId uuid.UUID,
	studentsId []uuid.UUID,
	contentsIDs []uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
) *Class {
	return &Class{
		id,
		name,
		teacherId,
		studentsId,
		contentsIDs,
		createdAt,
		updatedAt,
	}
}

func (c *Class) ID() uuid.UUID {
	return c.id
}

func (c *Class) Name() string {
	return c.name
}

func (c *Class) TeacherID() uuid.UUID {
	return c.teacherID
}

func (c *Class) StudentsID() []uuid.UUID {
	return c.studentsID
}

func (c *Class) ContentsID() []uuid.UUID {
	return c.contentsID
}

func (c *Class) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Class) UpdatedAt() time.Time {
	return c.updatedAt
}

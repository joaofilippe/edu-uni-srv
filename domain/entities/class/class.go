package classentities

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	id          uuid.UUID
	name        string
	teacherID   uuid.UUID
	studentsIDs []uuid.UUID
	contentsIDs []uuid.UUID
	createdAt   time.Time
	updatedAt   time.Time
}

func NewClass(
	id uuid.UUID,
	name string,
	teacherId uuid.UUID,
	studentsIds []uuid.UUID,
	contentsIDs []uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
) *Class {
	return &Class{
		id,
		name,
		teacherId,
		studentsIds,
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

func (c *Class) StudentsIDs() []uuid.UUID {
	return c.studentsIDs
}

func (c *Class) ContentsIDs() []uuid.UUID {
	return c.contentsIDs
}

func (c *Class) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Class) UpdatedAt() time.Time {
	return c.updatedAt
}

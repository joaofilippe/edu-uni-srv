package teacherentities

import (
	"time"

	"github.com/google/uuid"
)

type Teacher struct {
	id       uuid.UUID
	userID   uuid.UUID
	classesIDs []uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	active   bool
}

func NewTeacher(
	id, userID uuid.UUID,
	classesIDs []uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	active bool,
) *Teacher {
	return &Teacher{
		id,
		userID,
		classesIDs,
		createdAt,
		updatedAt,
		active,
	}
}

func (t *Teacher) ID() uuid.UUID {
	return t.id
}
func (t *Teacher) ClassesIDs() []uuid.UUID {
	return t.classesIDs
}

func (t *Teacher) UserID() uuid.UUID {
	return t.userID
}
 
func (t *Teacher) IsActive() bool {
	return t.active
}

func (t *Teacher) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Teacher) UpdatedAt() time.Time {
	return t.updatedAt
}
package teacherentities

import "github.com/google/uuid"

type Teacher struct {
	id       uuid.UUID
	userID   uuid.UUID
	name     string
	classesIDs []uuid.UUID
}

func NewTeacher(
	id, userID uuid.UUID,
	name string,
	classesIDs []uuid.UUID,
) *Teacher {
	return &Teacher{
		id,
		userID,
		name,
		classesIDs,
	}
}

func (t *Teacher) ID() uuid.UUID {
	return t.id
}

func (t *Teacher) Name() string {
	return t.name
}

func (t *Teacher) ClassIDs() []uuid.UUID {
	return t.classesIDs
}

func (t *Teacher) UserID() uuid.UUID {
	return t.userID
}

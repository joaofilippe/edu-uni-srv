package teacherentities

import "github.com/google/uuid"

type CreateTeacher struct {
	id       uuid.UUID
	userID   uuid.UUID
	name     string
	classesIDs []uuid.UUID
}

func NewCreateTeacher(
	userID uuid.UUID,
	name string,
	classesIDs []uuid.UUID,
) *CreateTeacher {
	return &CreateTeacher{
		userID:userID,
		name:name,
		classesIDs: classesIDs,
	}
}

func (t *CreateTeacher) ID() uuid.UUID {
	return t.id
}

func (t *CreateTeacher) UserID() uuid.UUID {
	return t.userID
}

func (t *CreateTeacher) Name() string {
	return t.name
}

func (t *CreateTeacher) ClassesIDs() []uuid.UUID {
	return t.classesIDs
}

func (t *CreateTeacher) EmptyID() bool {
	return len(t.id) == 0
}

func (t *CreateTeacher) SetID(id uuid.UUID) {
	t.id = id
}
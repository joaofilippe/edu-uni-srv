package teacherentities

import "github.com/google/uuid"

type CreateTeacher struct {
	id     uuid.UUID
	userID uuid.UUID
}

func NewCreateTeacher(
	userID uuid.UUID,
) *CreateTeacher {
	return &CreateTeacher{
		userID: userID,
	}
}

func (t *CreateTeacher) ID() uuid.UUID {
	return t.id
}

func (t *CreateTeacher) UserID() uuid.UUID {
	return t.userID
}

func (t *CreateTeacher) SetID(id uuid.UUID) {
	t.id = id
}

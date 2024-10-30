package teacher

import "github.com/google/uuid"

type Teacher struct {
	id       uuid.UUID
	userId   uuid.UUID
	name     string
	classIds []uuid.UUID
}

func NewTeacher(
	id, userId uuid.UUID,
	name string,
	classIds []uuid.UUID,
) *Teacher {
	return &Teacher{
		id,
		userId,
		name,
		classIds,
	}
}

func (t *Teacher) ID() uuid.UUID {
	return t.id
}

func (t *Teacher) Name() string {
	return t.name
}

func (t *Teacher) ClassIds() []uuid.UUID {
	return t.classIds
}

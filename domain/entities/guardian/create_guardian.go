package guardianentities

import "github.com/google/uuid"

type CreateGuardian struct {
	id        uuid.UUID
	userID    uuid.UUID
	studentID uuid.UUID
}

func NewCreateGuardian(
	userID uuid.UUID,
	studentID uuid.UUID,
) *CreateGuardian {
	return &CreateGuardian{
		userID:    userID,
		studentID: studentID,
	}
}

func (g *CreateGuardian) ID() uuid.UUID {
	return g.id
}

func (g *CreateGuardian) UserID() uuid.UUID {
	return g.userID
}

func (g *CreateGuardian) StudentID() uuid.UUID {
	return g.studentID
}

func (g *CreateGuardian) SetId(id uuid.UUID) {
	g.id = id
}

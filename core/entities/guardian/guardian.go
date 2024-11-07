package guardianentities

import (
	"time"

	"github.com/google/uuid"
)

type Guardian struct {
	id        uuid.UUID
	userID    uuid.UUID
	studentID uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	active    bool
}

func NewGuardian(
	id uuid.UUID,
	userID uuid.UUID,
	studentID uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	active bool,
) *Guardian {
	return &Guardian{
		id,
		userID,
		studentID,
		createdAt,
		updatedAt,
		active,
	}
}

func (g *Guardian) ID() uuid.UUID {
	return g.id
}

func (g *Guardian) UserID() uuid.UUID {
	return g.userID
}

func (g *Guardian) StudentID() uuid.UUID {
	return g.studentID
}

func (g *Guardian) CreatedAt() time.Time {
	return g.createdAt
}

func (g *Guardian) UpdatedAt() time.Time {
	return g.updatedAt
}

func (g *Guardian) Active() bool {
	return g.active
}

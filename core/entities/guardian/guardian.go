package guardianentities

import (
	"time"

	"github.com/google/uuid"
)

type Guardian struct {
	id        uuid.UUID
	userID    uuid.UUID
	studentID uuid.UUID
	name      string
	email     string
	password  string
	createdAt time.Time
	updatedAt time.Time
	active    bool
}

func NewGuardian(
	id uuid.UUID,
	userID uuid.UUID,
	name string,
	email string,
	password string,
	studentID uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	active bool,
) *Guardian {
	return &Guardian{
		id,
		userID,
		studentID,
		name,
		email,
		password,
		createdAt,
		updatedAt,
		active,
	}
}

func (g *Guardian) Id() uuid.UUID {
	return g.id
}

func (g *Guardian) UserID() uuid.UUID {
	return g.userID
}

func (g *Guardian) Name() string {
	return g.name
}

func (g *Guardian) Email() string {
	return g.email
}

func (g *Guardian) Password() string {
	return g.password
}

func (g *Guardian) StudentId() uuid.UUID {
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

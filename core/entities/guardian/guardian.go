package guardian

import "github.com/google/uuid"

type Guardian struct {
	id        uuid.UUID
	userID    uuid.UUID
	studentID uuid.UUID
	name      string
	email     string
	password  string
}

func NewGuardian(
	id uuid.UUID,
	userID uuid.UUID,
	name string,
	email string,
	password string,
	studentID uuid.UUID,
) *Guardian {
	return &Guardian{
		id,
		userID,
		studentID,
		name,
		email,
		password,
	}
}

func (g *Guardian) Id() uuid.UUID {
	return g.id
}

func (g *Guardian) UserId() uuid.UUID {
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

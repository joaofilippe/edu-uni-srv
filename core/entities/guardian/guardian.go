package guardian

import "github.com/google/uuid"

type Guardian struct {
	id        uuid.UUID
	userId    uuid.UUID
	name      string
	email     string
	password  string
	studentId uuid.UUID
}

func NewGuardian(
	id uuid.UUID,
	userId uuid.UUID,
	name string,
	email string,
	password string,
	studentId uuid.UUID,
) *Guardian {
	return &Guardian{
		id,
		userId,
		name,
		email,
		password,
		studentId,
	}
}

func (g *Guardian) Id() uuid.UUID {
	return g.id
}

func (g *Guardian) UserId() uuid.UUID {
	return g.userId
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
	return g.studentId
}

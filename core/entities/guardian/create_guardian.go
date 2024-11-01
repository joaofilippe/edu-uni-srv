package guardian

import "github.com/google/uuid"

type CreateGuardian struct {
	id        uuid.UUID
	userID    uuid.UUID
	studentID uuid.UUID
	name      string
	email     string
	password  string
}

func NewCreateGuardian(
	userID uuid.UUID,
	studentID uuid.UUID,
	name string,
	email string,
	password string,
) *CreateGuardian {
	return &CreateGuardian{
		userID:    userID,
		studentID: studentID,
		name:      name,
		email:     email,
		password:  password,
	}
}

func (g *CreateGuardian) UserID() uuid.UUID {
	return g.userID
}

func (g *CreateGuardian) StudentID() uuid.UUID {
	return g.studentID
}

func (g *CreateGuardian) Name() string {
	return g.name
}

func (g *CreateGuardian) Email() string {
	return g.email
}

func (g *CreateGuardian) Password() string {
	return g.password
}

func (g *CreateGuardian) SetId(id uuid.UUID) {
	g.id = id
}

func (g *CreateGuardian) EmptyID() bool {
	return len(g.id) == 0
}

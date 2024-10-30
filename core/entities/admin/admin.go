package admin

import "github.com/google/uuid"

type Admin struct {
	id       uuid.UUID
	userId   uuid.UUID
	name     string
	email    string
	password string
}

func NewAdmin(
	id uuid.UUID,
	userId uuid.UUID,
	name string,
	email string,
	password string,
) *Admin {
	return &Admin{
		id,
		userId,
		name,
		email,
		password,
	}
}

func (a *Admin) ID() uuid.UUID {
	return a.id
}

func (a *Admin) Name() string {
	return a.name
}

func (a *Admin) Email() string {
	return a.email
}

func (a *Admin) Password() string {
	return a.password
}

func (a *Admin) UserId() uuid.UUID {
	return a.userId
}

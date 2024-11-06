package adminentities

import (
	"time"

	"github.com/google/uuid"
)

type Admin struct {
	id        uuid.UUID
	userId    uuid.UUID
	name      string
	email     string
	createdAt time.Time
	updatedAt time.Time
	active    bool
}

func NewAdmin(
	id uuid.UUID,
	userId uuid.UUID,
	name string,
	email string,
	createdAt time.Time,
	updatedAt time.Time,
	active bool,
) *Admin {
	return &Admin{
		id,
		userId,
		name,
		email,
		createdAt,
		updatedAt,
		active,
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

func (a *Admin) UserID() uuid.UUID {
	return a.userId
}

func (a *Admin) CreatedAt() time.Time {
	return a.createdAt
}

func (a *Admin) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a *Admin) Active() bool {
	return a.active
}

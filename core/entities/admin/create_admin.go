package admin

import "github.com/google/uuid"

type CreateAdmin struct {
	id       uuid.UUID
	userID   uuid.UUID
	username string
	email    string
}

func NewCreateAdmin(
	userID uuid.UUID,
	username string,
	email string,
) *CreateAdmin {
	return &CreateAdmin{
		userID:   userID,
		username: username,
		email:    email,
	}
}

func (c *CreateAdmin) UserID() uuid.UUID {
	return c.userID
}

func (c *CreateAdmin) Username() string {
	return c.username
}

func (c *CreateAdmin) Email() string {
	return c.email
}

func (c *CreateAdmin) EmptyID() bool {
	return len(c.id) == 0
}

func (c *CreateAdmin) SetId(id uuid.UUID) {
	c.id = id
}

package admin

import "github.com/google/uuid"

type CreateAdmin struct {
	userId   uuid.UUID
	username string
	email    string
}

func NewCreateAdmin(
	userId uuid.UUID,
	username string,
	email string,
) *CreateAdmin {
	return &CreateAdmin{
		userId,
		username,
		email,
	}
}

func (c *CreateAdmin) UserId() uuid.UUID {
	return c.userId
}

func (c *CreateAdmin) Username() string {
	return c.username
}

func (c *CreateAdmin) Email() string {
	return c.email
}

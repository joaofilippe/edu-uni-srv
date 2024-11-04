package adminentities

import "github.com/google/uuid"

type CreateUseCase struct {
	id       uuid.UUID
	userID   uuid.UUID
	username string
	email    string
}

func NewCreateUseCase(
	userID uuid.UUID,
	username string,
	email string,
) *CreateUseCase {
	return &CreateUseCase{
		userID:   userID,
		username: username,
		email:    email,
	}
}

func (c *CreateUseCase) UserID() uuid.UUID {
	return c.userID
}

func (c *CreateUseCase) Username() string {
	return c.username
}

func (c *CreateUseCase) Email() string {
	return c.email
}

func (c *CreateUseCase) EmptyID() bool {
	return len(c.id) == 0
}

func (c *CreateUseCase) SetId(id uuid.UUID) {
	c.id = id
}

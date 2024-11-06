package adminentities

import "github.com/google/uuid"

type Create struct {
	id     uuid.UUID
	userID uuid.UUID
}

func NewCreate(
	userID uuid.UUID,
) *Create {
	return &Create{
		id:     uuid.UUID{},
		userID: userID,
	}
}

func (c *Create) ID() uuid.UUID {
	return c.id
}

func (c *Create) UserID() uuid.UUID {
	return c.userID
}

func (c *Create) SetId(id uuid.UUID) {
	c.id = id
}

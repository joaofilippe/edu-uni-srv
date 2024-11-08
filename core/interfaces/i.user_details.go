package interfaces

import "github.com/google/uuid"

type IUserDetails interface {
	ID() uuid.UUID
	UserID() uuid.UUID
}

package interfaces

import "github.com/google/uuid"

type IUserDetails interface {
	UserID() uuid.UUID
}
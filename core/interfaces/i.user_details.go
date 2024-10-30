package interfaces

import "github.com/google/uuid"

type IUserDetails interface {
	UserId() uuid.UUID
}
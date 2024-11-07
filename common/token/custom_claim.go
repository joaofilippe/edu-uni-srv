package token

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/interfaces"
)

type ICustomClaim interface {
	ID() uuid.UUID
	UserDetails() interfaces.IUserDetails
	UserType() enums.UserType
}
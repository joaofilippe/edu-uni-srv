package tokenmanager

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/enums"
	"github.com/joaofilippe/edu-uni-srv/domain/interfaces"
)

type ICustomClaim interface {
	ID() uuid.UUID
	UserDetails() interfaces.IUserDetails
	UserType() enums.UserType
}

package user

import (
	"github.com/google/uuid"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
)

type IUserRepo interface {
	Save(user *userEntities.User) (uuid.UUID,error)
}
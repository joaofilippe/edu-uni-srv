package repositories

import (
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
)

type IUserRepo interface {
	Save(user *userEntities.User) error
}

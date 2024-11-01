package repositories

import (
	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type IAdminRepo interface {
	Save(admin *adminEntities.CreateAdmin) error
}

package repositories

import (
	"github.com/google/uuid"

	adminEntities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
)

type IAdminRepo interface {
	Save(admin *adminEntities.Admin) (uuid.UUID, error)
}

package repositories

import guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"

type IGuardianRepo interface {
	Save(guardian *guardianEntities.CreateGuardian) error
}

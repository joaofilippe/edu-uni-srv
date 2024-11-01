package repositories

import guardianEntities "github.com/joaofilippe/edu-uni-srv/core/entities/guardian"

type GuardianRepository interface {
	Save(guardian *guardianEntities.CreateGuardian) error
}
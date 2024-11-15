package iservices

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/class"
)

type IClassService interface {
	Create(class *classentities.CreateClass) (uuid.UUID, error)
}
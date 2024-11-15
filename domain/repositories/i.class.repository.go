package irepositories

import (
	"github.com/joaofilippe/edu-uni-srv/domain/entities/class"
)

type IClassRepo interface {
	Save(class *classentities.CreateClass)  error
}
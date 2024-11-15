package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/class"
	"github.com/joaofilippe/edu-uni-srv/domain/usecases/class"
)

type ClassService struct {
	createUseCase *classusecases.CreateUseCase
}

func NewClassService(createUseCase *classusecases.CreateUseCase) *ClassService {
	return &ClassService{createUseCase: createUseCase}
}

func (cs *ClassService) Create(createClass *classentities.CreateClass) (uuid.UUID, error) {
	return cs.createUseCase.Execute(createClass)
}

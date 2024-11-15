package classusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/class"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type CreateUseCase struct {
	classRepository irepositories.IClassRepo
}

func NewCreateUseCase(classRepository irepositories.IClassRepo) *CreateUseCase {
	return &CreateUseCase{
		classRepository: classRepository,
	}
}

func (c *CreateUseCase) Execute(class *classentities.CreateClass) (uuid.UUID, error) {
	class.SetID(uuid.New())
	return class.ID(), c.classRepository.Save(class)
}

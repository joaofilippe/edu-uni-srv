package contentusecases

import (
	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type CreateUseCase struct {
	contentRepository irepositories.IContentRepo
}

func NewCreateUseCase(
	contentRepository *irepositories.IContentRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*contentRepository,
	}
}

func (c *CreateUseCase) Execute(
	content *contententities.CreateContent,
) (uuid.UUID, error) {
	content.SetID(uuid.New())
	return content.ID(), c.contentRepository.Save(content)
}

package contentusecases

import (
	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type CreateContentUseCase struct {
	contentRepository irepositories.IContentRepo
}

func NewCreateContentUseCase(
	contentRepository *irepositories.IContentRepo,
) *CreateContentUseCase {
	return &CreateContentUseCase{
		*contentRepository,
	}
}

func (c *CreateContentUseCase) Execute(
	content *contententities.CreateContent,
) error {
	content.SetID(uuid.New())
	return c.contentRepository.Save(content)
}

package services

import (
	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	contentusecases "github.com/joaofilippe/edu-uni-srv/domain/usecases/content"
)

type ContentService struct {
	createUseCase *contentusecases.CreateContentUseCase
}

func NewContentService(
	createUseCase *contentusecases.CreateContentUseCase,
) *ContentService {
	return &ContentService{
		createUseCase,
	}
}

func (cs *ContentService) Create(createContent *contententities.CreateContent) (uuid.UUID,error) {
	return cs.createUseCase.Execute(createContent)
}
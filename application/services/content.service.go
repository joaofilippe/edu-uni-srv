package services

import (
	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
	contentusecases "github.com/joaofilippe/edu-uni-srv/domain/usecases/content"
)

type ContentService struct {
	createUseCase *contentusecases.CreateUseCase
}

func NewContentService(
	createUseCase *contentusecases.CreateUseCase,
) *ContentService {
	return &ContentService{
		createUseCase,
	}
}

func (cs *ContentService) Create(createContent *contententities.CreateContent) (uuid.UUID, error) {
	return cs.createUseCase.Execute(createContent)
}

func (cs *ContentService) FindByID(id uuid.UUID) (*contententities.Content, error) {
	return nil, nil
}
func (cs *ContentService) FindByStudentID(studentID uuid.UUID) ([]*contententities.Content, error) {
	return nil, nil
}

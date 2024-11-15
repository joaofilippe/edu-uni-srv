package iservices

import (
	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
)

type IContentService interface {
	Create(content *contententities.CreateContent) (uuid.UUID, error)
	FindByID(id uuid.UUID) (*contententities.Content, error)
	FindByStudentID(studentID uuid.UUID) ([]*contententities.Content, error)
}

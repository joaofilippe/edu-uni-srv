package irepositories

import (
	"github.com/google/uuid"
	contententities "github.com/joaofilippe/edu-uni-srv/domain/entities/content"
)

type IContentRepo interface {
	Save(content *contententities.CreateContent) error
	FindByID(id uuid.UUID) (*contententities.Content, error)
	FindByStudentID(studentID uuid.UUID) ([]*contententities.Content, error)
}

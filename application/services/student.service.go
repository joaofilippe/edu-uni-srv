package services

import (
	"github.com/google/uuid"
	studententities "github.com/joaofilippe/edu-uni-srv/core/entities/student"
	studentusecase "github.com/joaofilippe/edu-uni-srv/core/usecases/student"
)

type StudentService struct {
	createUseCase       *studentusecase.CreateUseCase
	findAllUseCase      *studentusecase.FindAllUseCase
	findByIDUseCase     *studentusecase.FindByIDUseCase
	findByUserIDUseCase *studentusecase.FindByUserIDUseCase
	updateUseCase       *studentusecase.UpdateUseCase
	deleteUseCase       *studentusecase.DeleteUseCase
}

func NewStudentService(
	createUseCase *studentusecase.CreateUseCase,
	findAllUseCase *studentusecase.FindAllUseCase,
	findByIDUseCase *studentusecase.FindByIDUseCase,
	findByUserIDUseCase *studentusecase.FindByUserIDUseCase,
	updateUseCase *studentusecase.UpdateUseCase,
	deleteUseCase *studentusecase.DeleteUseCase,
) *StudentService {
	return &StudentService{
		createUseCase:       createUseCase,
		findAllUseCase:      findAllUseCase,
		findByIDUseCase:     findByIDUseCase,
		findByUserIDUseCase: findByUserIDUseCase,
		updateUseCase:       updateUseCase,
		deleteUseCase:       deleteUseCase,
	}
}

func (s *StudentService) Create(student *studententities.Student) error {
	return nil
}

func (s *StudentService) FindAll() ([]*studententities.Student, error) {
	return []*studententities.Student{}, nil
}

func (s *StudentService) FindById(id uuid.UUID) (*studententities.Student, error) {
	return &studententities.Student{}, nil
}

func (s *StudentService) FindByUserID(userID uuid.UUID) (*studententities.Student, error) {
	return &studententities.Student{}, nil
}

func (s *StudentService) Update(student *studententities.Student) error {
	return nil
}

func (s *StudentService) Delete(id string) error {
	return nil
}

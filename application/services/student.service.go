package services

import (
	"github.com/google/uuid"
	studententities "github.com/joaofilippe/edu-uni-srv/domain/entities/student"
	studentusecase "github.com/joaofilippe/edu-uni-srv/domain/usecases/student"
)

type StudentService struct {
	createUseCase       *studentusecase.CreateUseCase
	enrollUseCase       *studentusecase.EnrollUseCase
	findAllUseCase      *studentusecase.FindAllUseCase
	findByIDUseCase     *studentusecase.FindByIDUseCase
	findByUserIDUseCase *studentusecase.FindByUserIDUseCase
	updateUseCase       *studentusecase.UpdateUseCase
	deleteUseCase       *studentusecase.DeleteUseCase
}

func NewStudentService(
	createUseCase *studentusecase.CreateUseCase,
	enrollUseCase *studentusecase.EnrollUseCase,
	findAllUseCase *studentusecase.FindAllUseCase,
	findByIDUseCase *studentusecase.FindByIDUseCase,
	findByUserIDUseCase *studentusecase.FindByUserIDUseCase,
	updateUseCase *studentusecase.UpdateUseCase,
	deleteUseCase *studentusecase.DeleteUseCase,
) *StudentService {
	return &StudentService{
		createUseCase,
		enrollUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByUserIDUseCase,
		updateUseCase,
		deleteUseCase,
	}
}

func (s *StudentService) Create(student *studententities.CreateStudent) (uuid.UUID, error) {
	return s.createUseCase.Execute(student)
}

func (s *StudentService) Enroll(studentID, classID uuid.UUID) error {
	return s.enrollUseCase.Execute(studentID, classID)
}

func (s *StudentService) FindAll() ([]*studententities.Student, error) {
	return []*studententities.Student{}, nil
}

func (s *StudentService) FindById(id uuid.UUID) (*studententities.Student, error) {
	return &studententities.Student{}, nil
}

func (s *StudentService) FindByUserID(userID uuid.UUID) (*studententities.Student, error) {
	return s.findByIDUseCase.Execute(userID)
}

func (s *StudentService) Update(student *studententities.Student) error {
	return nil
}

func (s *StudentService) Delete(id string) error {
	return nil
}

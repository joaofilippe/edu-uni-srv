package services

import (
	"github.com/google/uuid"

	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	teacherusecase "github.com/joaofilippe/edu-uni-srv/domain/usecases/teacher"
)

type TeacherService struct {
	createUseCase       *teacherusecase.CreateUseCase
	findAllUseCase      *teacherusecase.FindAllUseCase
	findByIDUseCase     *teacherusecase.FindByIDUseCase
	findByUserIDUseCase *teacherusecase.FindByUserIDUseCase
	updateUseCase       *teacherusecase.UpdateUseCase
	deleteUseCase       *teacherusecase.DeleteUseCase
}

func NewTeacherService(
	createUseCase *teacherusecase.CreateUseCase,
	findAllUseCase *teacherusecase.FindAllUseCase,
	findByIDUseCase *teacherusecase.FindByIDUseCase,
	findByUserIDUseCase *teacherusecase.FindByUserIDUseCase,
	updateUseCase *teacherusecase.UpdateUseCase,
	deleteUseCase *teacherusecase.DeleteUseCase,
) *TeacherService {
	return &TeacherService{
		createUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByUserIDUseCase,
		updateUseCase,
		deleteUseCase,
	}
}

func (ts *TeacherService) Create(teacher *teacherentities.Teacher) error {
	return nil
}

func (ts *TeacherService) FindAll() ([]*teacherentities.Teacher, error) {
	return []*teacherentities.Teacher{}, nil
}

func (ts *TeacherService) FindById(id uuid.UUID) (*teacherentities.Teacher, error) {
	return &teacherentities.Teacher{}, nil
}

func (ts *TeacherService) FindByUserID(userID uuid.UUID) (*teacherentities.Teacher, error) {
	return &teacherentities.Teacher{}, nil
}

func (ts *TeacherService) Update(teacher *teacherentities.Teacher) error {
	return nil
}

func (ts *TeacherService) Delete(id uuid.UUID) error {
	return nil
}

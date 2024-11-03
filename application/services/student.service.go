package services

import (
	studentUseCases "github.com/joaofilippe/edu-uni-srv/core/usecases/student"
)

type StudentService struct {
	createUseCase       *studentUseCases.CreateUseCase
	findAllUseCase      *studentUseCases.FindAllUsecase
	findByIDUseCase     *studentUseCases.FindByIDUseCase
	findByUserIDUseCase *studentUseCases.FindByUserIDUseCase
	updateUseCase       *studentUseCases.UpdateUseCase
	deleteUseCase       *studentUseCases.DeleteUseCase
}

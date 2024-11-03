package services

import (
	"github.com/joaofilippe/edu-uni-srv/core/usecases/student"
)

type StudentService struct {
	createUseCase       *studentusecase.CreateUseCase
	findAllUseCase      *studentusecase.FindAllUsecase
	findByIDUseCase     *studentusecase.FindByIDUseCase
	findByUserIDUseCase *studentusecase.FindByUserIDUseCase
	updateUseCase       *studentusecase.UpdateUseCase
	deleteUseCase       *studentusecase.DeleteUseCase
}

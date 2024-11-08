package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/guardian"
	"github.com/joaofilippe/edu-uni-srv/domain/usecases/guardian"
)

type GuardianService struct {
	createUseCase       *guardianusecases.CreateUseCase
	findAllUseCase      *guardianusecases.FindAllUseCase
	findByIDUseCase     *guardianusecases.FindByIDUseCase
	findByUserIDUseCase *guardianusecases.FindByUserIDUseCase
	updateUseCase       *guardianusecases.UpdateUseCase
	deleteUseCase       *guardianusecases.DeleteUseCase
}

func NewGuardianService(
	createUseCase *guardianusecases.CreateUseCase,
	findAllUseCase *guardianusecases.FindAllUseCase,
	findByIDUseCase *guardianusecases.FindByIDUseCase,
	findByUserIDUseCase *guardianusecases.FindByUserIDUseCase,
	updateUseCase *guardianusecases.UpdateUseCase,
	deleteUseCase *guardianusecases.DeleteUseCase,
) *GuardianService {
	return &GuardianService{
		createUseCase:       createUseCase,
		findAllUseCase:      findAllUseCase,
		findByIDUseCase:     findByIDUseCase,
		findByUserIDUseCase: findByUserIDUseCase,
		updateUseCase:       updateUseCase,
		deleteUseCase:       deleteUseCase,
	}
}

func (gs *GuardianService) Create(guardian *guardianentities.Guardian) error {
	return nil
}
func (gs *GuardianService) FindAll() ([]*guardianentities.Guardian, error) {
	return nil, nil
}
func (gs *GuardianService) FindById(id uuid.UUID) (*guardianentities.Guardian, error) {
	return nil, nil
}
func (gs *GuardianService) FindByUserID(userID uuid.UUID) (*guardianentities.Guardian, error) {
	return nil, nil
}
func (gs *GuardianService) Update(guardian *guardianentities.Guardian) error {
	return nil
}
func (gs *GuardianService) Delete(id uuid.UUID) error {
	return nil
}

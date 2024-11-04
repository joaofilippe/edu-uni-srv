package services

import (
	"github.com/google/uuid"

	adminentities "github.com/joaofilippe/edu-uni-srv/core/entities/admin"
	adminusecases "github.com/joaofilippe/edu-uni-srv/core/usecases/admin"
)

type AdminService struct {
	createUseCase       *adminusecases.CreateUseCase
	findAllUseCase      *adminusecases.FindAllUseCase
	findByIDUseCase     *adminusecases.FindByIDUseCase
	findByUserIDUseCase *adminusecases.FindByUserIDUseCase
	updateUseCase       *adminusecases.UpdateUseCase
	deleteUseCase       *adminusecases.DeleteUseCase
}

func NewAdminService(
	createUseCase *adminusecases.CreateUseCase,
	findAllUseCase *adminusecases.FindAllUseCase,
	findByIDUseCase *adminusecases.FindByIDUseCase,
	findByUserIDUseCase *adminusecases.FindByUserIDUseCase,
	updateUseCase *adminusecases.UpdateUseCase,
	deleteUseCase *adminusecases.DeleteUseCase,
) *AdminService {
	return &AdminService{
		createUseCase:       createUseCase,
		findAllUseCase:      findAllUseCase,
		findByIDUseCase:     findByIDUseCase,
		findByUserIDUseCase: findByUserIDUseCase,
		updateUseCase:       updateUseCase,
		deleteUseCase:       deleteUseCase,
	}
}

func (as *AdminService) Create(admin *adminentities.Admin) error {
	return nil
}
func (as *AdminService) FindAll() ([]*adminentities.Admin, error) {
	return []*adminentities.Admin{}, nil
}
func (as *AdminService) FindById(id uuid.UUID) (*adminentities.Admin, error) {
	return &adminentities.Admin{}, nil
}
func (as *AdminService) FindByUserID(userID uuid.UUID) (*adminentities.Admin, error) {
	return &adminentities.Admin{}, nil
}
func (as *AdminService) Update(admin *adminentities.Admin) error {
	return nil
}
func (as *AdminService) Delete(id uuid.UUID) error {
	return nil
}

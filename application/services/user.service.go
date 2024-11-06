package services

import (
	"github.com/google/uuid"

	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	userUseCases "github.com/joaofilippe/edu-uni-srv/core/usecases/user"
)

type UserService struct {
	createUseCase      *userUseCases.CreateUseCase
	findAllUseCase     *userUseCases.FindAllUseCase
	findByIDUseCase    *userUseCases.FindByIDUseCase
	findByEmailUseCase *userUseCases.FindByEmailUseCase
	updateUseCase      *userUseCases.UpdateUseCase
	deleteUseCase      *userUseCases.DeleteUseCase
}

func NewUserService(
	createUseCase *userUseCases.CreateUseCase,
	findAllUseCase *userUseCases.FindAllUseCase,
	findByIDUseCase *userUseCases.FindByIDUseCase,
	findByEmailUseCase *userUseCases.FindByEmailUseCase,
	updateUseCase *userUseCases.UpdateUseCase,
	deleteUseCase *userUseCases.DeleteUseCase,
) *UserService {
	return &UserService{
		createUseCase:      createUseCase,
		findAllUseCase:     findAllUseCase,
		findByIDUseCase:    findByIDUseCase,
		findByEmailUseCase: findByEmailUseCase,
		updateUseCase:      updateUseCase,
		deleteUseCase:      deleteUseCase,
	}
}

func (u *UserService) Create(user *userEntities.CreateUser) (uuid.UUID, error) {
	return u.createUseCase.Execute(user)
}
func (u *UserService) FindAll() ([]*userEntities.User, error) {
	return []*userEntities.User{}, nil
}
func (u *UserService) FindById(id uuid.UUID) (*userEntities.User, error) {
	return &userEntities.User{}, nil
}
func (u *UserService) FindByEmail(email string) (*userEntities.User, error) {
	return &userEntities.User{}, nil
}
func (u *UserService) Update(user *userEntities.User) error {
	return nil
}
func (u *UserService) Delete(id string) error {
	return nil
}

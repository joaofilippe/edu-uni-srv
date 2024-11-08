package services

import (
	"github.com/google/uuid"

	userEntities "github.com/joaofilippe/edu-uni-srv/domain/entities/user"
	userUseCases "github.com/joaofilippe/edu-uni-srv/domain/usecases/user"
)

type UserService struct {
	createUseCase      *userUseCases.CreateUseCase
	loginUseCase       *userUseCases.LoginUseCase
	findAllUseCase     *userUseCases.FindAllUseCase
	findByIDUseCase    *userUseCases.FindByIDUseCase
	findByEmailUseCase *userUseCases.FindByEmailUseCase
	updateUseCase      *userUseCases.UpdateUseCase
	deleteUseCase      *userUseCases.DeleteUseCase
}

func NewUserService(
	createUseCase *userUseCases.CreateUseCase,
	loginUseCase *userUseCases.LoginUseCase,
	findAllUseCase *userUseCases.FindAllUseCase,
	findByIDUseCase *userUseCases.FindByIDUseCase,
	findByEmailUseCase *userUseCases.FindByEmailUseCase,
	updateUseCase *userUseCases.UpdateUseCase,
	deleteUseCase *userUseCases.DeleteUseCase,
) *UserService {
	return &UserService{
		createUseCase,
		loginUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByEmailUseCase,
		updateUseCase,
		deleteUseCase,
	}
}

func (u *UserService) Create(user *userEntities.CreateUser) (uuid.UUID, error) {
	return u.createUseCase.Execute(user)
}

func (u *UserService) Login(email, password string) (string, error) {
	return u.loginUseCase.Execute(email, password)
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

package services

import (
	"github.com/google/uuid"

	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	userUseCases "github.com/joaofilippe/edu-uni-srv/core/usecases/user"
)

type UserService struct {
	CreateUseCase userUseCases.CreateUseCase
	FindAllUseCase 	 userUseCases.FindAllUseCase
	FindByIDUseCase 	 userUseCases.FindByIDUseCase
	FindByEmailUseCase  userUseCases.FindByEmailUseCase
	UpdateUseCase userUseCases.UpdateUseCase
	DeleteUseCase userUseCases.DeleteUseCase
}

func (u *UserService) Create(user *userEntities.User) error{
	return nil
}
func (u *UserService) FindAll() ([]*userEntities.User, error){
	return []*userEntities.User{},nil
}
func (u *UserService) FindById(id uuid.UUID) (*userEntities.User, error){
	return &userEntities.User{} ,nil
}
func (u *UserService) FindByEmail(email string) (*userEntities.User, error){
	return &userEntities.User{}, nil
}
func (u *UserService) FindByUsername(username string) (*userEntities.User, error){
	return &userEntities.User{}, nil
}
func (u *UserService) Update(user *userEntities.User) error{
	return nil
}
func (u *UserService) Delete(id string) error{
	return nil
}

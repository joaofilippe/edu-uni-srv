package user

import (
	"github.com/google/uuid"
	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	userRepository "github.com/joaofilippe/edu-uni-srv/core/repositories/user"
)

type CreateUserUseCase struct {
	userRepository userRepository.IUserRepo
}

func NewCreateUserUseCase(userRepository userRepository.IUserRepo) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(createUser *userEntities.CreateUser) (uuid.UUID, error){
	id := uuid.New()
	user, error := userEntities.NewUser(
		id, 
		createUser.Username(), 
		createUser.Password(), 
		createUser.Email(), 
		createUser.UserType(), 
		createUser.UserDetails(),
	)
	if error != nil {
		return uuid.UUID{}, error
	}
	
	return uc.userRepository.Save(user)
} 
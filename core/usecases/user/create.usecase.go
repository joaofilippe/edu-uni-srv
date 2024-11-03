package user

import (
	"github.com/google/uuid"

	userEntities "github.com/joaofilippe/edu-uni-srv/core/entities/user"
	userRepository "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	userRepository userRepository.IUserRepo
}

func NewCreateUseCase(userRepository *userRepository.IUserRepo) *CreateUseCase {
	return &CreateUseCase{
		userRepository: *userRepository,
	}
}

func (uc *CreateUseCase) Execute(createUser *userEntities.CreateUser) (uuid.UUID, error) {
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

	err := uc.userRepository.Save(user)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

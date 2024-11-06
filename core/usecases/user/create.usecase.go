package userusecases

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/common/errors"
	"github.com/joaofilippe/edu-uni-srv/core/entities/user"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	userRepository irepositories.IUserRepo
}

func NewCreateUseCase(userRepository *irepositories.IUserRepo) *CreateUseCase {
	return &CreateUseCase{
		userRepository: *userRepository,
	}
}

func (uc *CreateUseCase) Execute(createUser *userentities.CreateUser) (uuid.UUID, error) {
	if createUser.EmptyID() {
		createUser.SetID(uuid.New())
	}

	if createUser.Username() == "" {
		return uuid.UUID{}, usecaseerrors.ErrUserNoUsername
	}

	if createUser.Password() == "" {
		return uuid.UUID{}, usecaseerrors.ErrUserNoPassword
	}

	if createUser.Email() == "" {
		return uuid.UUID{}, usecaseerrors.ErrUserNoEmail
	}

	if !createUser.ValidateEmail() {
		return uuid.UUID{}, usecaseerrors.ErrUserInvalidEmail
	}

	err := uc.userRepository.Save(createUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return createUser.ID(), err
}

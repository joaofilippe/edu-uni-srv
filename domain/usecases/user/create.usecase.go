package userusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/edu-uni-srv/common/errors"
	"github.com/joaofilippe/edu-uni-srv/common/password"
	"github.com/joaofilippe/edu-uni-srv/domain/entities/user"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
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
		return uuid.UUID{}, errors.ErrUserNoUsername
	}

	if createUser.Password() == "" {
		return uuid.UUID{}, errors.ErrUserNoPassword
	}

	u, err := uc.userRepository.FindByEmail(createUser.Email())
	if err != nil {
		return uuid.UUID{}, err
	}

	if u != nil {
		return uuid.UUID{}, errors.ErrUserEmailAlreadyExists
	}

	hashedPassword, err := password.HashPassword(createUser.Password())
	if err != nil {
		return uuid.UUID{}, err
	}

	createUser.SetPassword(hashedPassword)

	if createUser.Email() == "" {
		return uuid.UUID{}, errors.ErrUserNoEmail
	}

	if !createUser.ValidateEmail() {
		return uuid.UUID{}, errors.ErrUserInvalidEmail
	}

	err = uc.userRepository.Save(createUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return createUser.ID(), err
}

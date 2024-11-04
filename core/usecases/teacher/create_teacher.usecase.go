package teacherusecase

import (
	"github.com/google/uuid"

	teacherentities "github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateUseCase struct {
	teacherRepo irepositories.ITeacherRepo
	userRepo    irepositories.IUserRepo
}

func NewCreateUseCase(
	teacherRepo irepositories.ITeacherRepo,
	userRepo irepositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		teacherRepo,
		userRepo,
	}
}

func (uc *CreateUseCase) Execute(teacher *teacherentities.CreateTeacher) (uuid.UUID, error) {
	if teacher.EmptyID() {
		teacher.SetID(uuid.New())
	}

	err := uc.teacherRepo.Save(teacher)
	if err != nil {
		return uuid.UUID{}, err
	}

	user, err := uc.userRepo.FindByID(teacher.UserID())
	if err != nil {
		return uuid.UUID{}, err
	}

	teacherType := enums.Teacher

	newUser := user.CopyWith(
		nil,
		nil,
		nil,
		&teacherType,
		nil,
	)

	err = uc.userRepo.Update(newUser)
	if err != nil {
		return uuid.UUID{}, err
	}

	return teacher.ID(), nil
}

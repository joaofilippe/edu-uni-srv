package teacherusecase

import (
	"github.com/google/uuid"

	"github.com/joaofilippe/edu-uni-srv/core/entities/teacher"
	"github.com/joaofilippe/edu-uni-srv/core/enums"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
)

type CreateTeacherUseCase struct {
	teacherRepo irepositories.ITeacherRepo
	userRepo    irepositories.IUserRepo
}

func NewCreateTeacherUseCase(
	teacherRepo irepositories.ITeacherRepo,
	userRepo irepositories.IUserRepo,
) *CreateTeacherUseCase {
	return &CreateTeacherUseCase{
		teacherRepo,
		userRepo,
	}
}

func (uc *CreateTeacherUseCase) Execute(teacher *teacherentities.CreateTeacher) (uuid.UUID, error) {
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

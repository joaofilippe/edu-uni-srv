package teacherusecases

import (
	"github.com/google/uuid"

	teacherentities "github.com/joaofilippe/edu-uni-srv/domain/entities/teacher"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
)

type CreateUseCase struct {
	teacherRepo irepositories.ITeacherRepo
	userRepo    irepositories.IUserRepo
}

func NewCreateUseCase(
	teacherRepo *irepositories.ITeacherRepo,
	userRepo *irepositories.IUserRepo,
) *CreateUseCase {
	return &CreateUseCase{
		*teacherRepo,
		*userRepo,
	}
}

func (uc *CreateUseCase) Execute(teacher *teacherentities.CreateTeacher) (uuid.UUID, error) {
	teacher.SetID(uuid.New())

	err := uc.teacherRepo.Save(teacher)
	if err != nil {
		err = uc.userRepo.Delete(teacher.UserID())
		if err != nil {
			return uuid.UUID{}, err
		}
		return uuid.UUID{}, err
	}

	return teacher.ID(), nil
}

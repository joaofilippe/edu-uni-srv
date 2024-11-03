package studentusecase

import "github.com/joaofilippe/edu-uni-srv/core/repositories"

type DeleteUseCase struct {
	StudentRepo repositories.IStudentRepo
}

func NewDeleteUseCase(studentRepo *repositories.IStudentRepo) *DeleteUseCase {
	return &DeleteUseCase{
		StudentRepo: *studentRepo,
	}
}

func (u *DeleteUseCase) Execute() error {
	return nil
}

package studentusecases

import "github.com/joaofilippe/edu-uni-srv/domain/repositories"

type DeleteUseCase struct {
	StudentRepo irepositories.IStudentRepo
}

func NewDeleteUseCase(studentRepo *irepositories.IStudentRepo) *DeleteUseCase {
	return &DeleteUseCase{
		StudentRepo: *studentRepo,
	}
}

func (u *DeleteUseCase) Execute() error {
	return nil
}

package studentusecases

import "github.com/joaofilippe/edu-uni-srv/domain/repositories"

type UpdateUseCase struct {
	StudentRepo irepositories.IStudentRepo
}

func NewUpdateUseCase(studentRepo *irepositories.IStudentRepo) *UpdateUseCase {
	return &UpdateUseCase{
		StudentRepo: *studentRepo,
	}
}

func (u *UpdateUseCase) Execute() error {
	return nil
}

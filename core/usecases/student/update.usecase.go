package studentusecase

import "github.com/joaofilippe/edu-uni-srv/core/repositories"

type UpdateUseCase struct {
	StudentRepo repositories.IStudentRepo
}

func NewUpdateUseCase(studentRepo *repositories.IStudentRepo) *UpdateUseCase {
	return &UpdateUseCase{
		StudentRepo: *studentRepo,
	}
}


func (u *UpdateUseCase) Execute() error {
	return nil
}
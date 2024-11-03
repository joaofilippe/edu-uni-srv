package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	"github.com/joaofilippe/edu-uni-srv/core/repositories"
	studentusecase "github.com/joaofilippe/edu-uni-srv/core/usecases/student"
)

func StudentServiceFactory(
	studentRepository *repositories.IStudentRepo,
	userRepository *repositories.IUserRepo,
) *services.StudentService {
	createUseCase := studentusecase.NewCreateUseCase(
		studentRepository,
		userRepository,
	)
	findAllUseCase := studentusecase.NewFindAllUseCase(studentRepository)
	findByID := studentusecase.NewFindByIDUseCase(studentRepository)
	findByUserID := studentusecase.NewFindByUserIDUseCase(studentRepository)
	update := studentusecase.NewUpdateUseCase(studentRepository)
	delete := studentusecase.NewDeleteUseCase(studentRepository)

	return services.NewStudentService(
		createUseCase,
		findAllUseCase,
		findByID,
		findByUserID,
		update,
		delete,
	)
}

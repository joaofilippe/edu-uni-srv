package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
	studentusecase "github.com/joaofilippe/edu-uni-srv/domain/usecases/student"
)

func StudentServiceFactory(
	studentRepository irepositories.IStudentRepo,
	userRepository irepositories.IUserRepo,
) *services.StudentService {
	createUseCase := studentusecase.NewCreateUseCase(
		&studentRepository,
		&userRepository,
	)
	findAllUseCase := studentusecase.NewFindAllUseCase(&studentRepository)
	findByID := studentusecase.NewFindByIDUseCase(&studentRepository)
	findByUserID := studentusecase.NewFindByUserIDUseCase(&studentRepository)
	update := studentusecase.NewUpdateUseCase(&studentRepository)
	deleteUsecase := studentusecase.NewDeleteUseCase(&studentRepository)

	return services.NewStudentService(
		createUseCase,
		findAllUseCase,
		findByID,
		findByUserID,
		update,
		deleteUsecase,
	)
}

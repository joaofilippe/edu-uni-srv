package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
	teacherusecase "github.com/joaofilippe/edu-uni-srv/core/usecases/teacher"
)

func TeacherServiceFactory(
	teacherRepository *irepositories.ITeacherRepo,
	userRepository *irepositories.IUserRepo,
) *services.TeacherService {
	createUseCase := teacherusecase.NewCreateUseCase(teacherRepository, userRepository)
	findAllUseCase := teacherusecase.NewFindAllUseCase(teacherRepository)
	findByIDUseCase := teacherusecase.NewFindByIDUseCase(teacherRepository)
	findByUserIDUseCase := teacherusecase.NewFindByUserIDUseCase(teacherRepository)
	updateUseCase := teacherusecase.NewUpdateUseCase(teacherRepository)
	deleteUseCase := teacherusecase.NewDeleteUseCase(teacherRepository)

	return services.NewTeacherService(
		createUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByUserIDUseCase,
		updateUseCase,
		deleteUseCase,
	)
	

}

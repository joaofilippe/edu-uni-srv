package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
	guardianusecases "github.com/joaofilippe/edu-uni-srv/domain/usecases/guardian"
)

func GuardianServiceFactory(
	guardianRepository irepositories.IGuardianRepo,
	studentRepository irepositories.IStudentRepo,
	userRepository irepositories.IUserRepo,
) *services.GuardianService {
	createUseCase := guardianusecases.NewCreateUseCase(&guardianRepository, &studentRepository,&userRepository)
	findAllUseCase := guardianusecases.NewFindAllUseCase(&guardianRepository)
	findByIDUseCase := guardianusecases.NewFindByIDUseCase(&guardianRepository)
	findByUserIDUseCase := guardianusecases.NewFindByUserIDUseCase(&guardianRepository)
	updateUseCase := guardianusecases.NewUpdateUseCase(&guardianRepository)
	deleteUseCase := guardianusecases.NewDeleteUseCase(&guardianRepository)

	return services.NewGuardianService(
		createUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByUserIDUseCase,
		updateUseCase,
		deleteUseCase,
	)

}

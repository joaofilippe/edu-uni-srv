package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	irepositories "github.com/joaofilippe/edu-uni-srv/core/repositories"
	guardianusecases "github.com/joaofilippe/edu-uni-srv/core/usecases/guardian"
)

func GuardianServiceFactory(
	guardianRepository irepositories.IGuardianRepo,
	userRepository irepositories.IUserRepo,
) *services.GuardianService {
	createUseCase := guardianusecases.NewCreateUseCase(&guardianRepository, &userRepository)
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

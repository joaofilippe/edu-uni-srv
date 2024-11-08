package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	irepositories "github.com/joaofilippe/edu-uni-srv/domain/repositories"
	adminusecases "github.com/joaofilippe/edu-uni-srv/domain/usecases/admin"
)

func AdminServiceFactory(
	adminRepository irepositories.IAdminRepo,
	userRepository irepositories.IUserRepo,
) *services.AdminService {
	createUseCase := adminusecases.NewCreateUseCase(&adminRepository, &userRepository)
	findAllUseCase := adminusecases.NewFindAllUseCase(&adminRepository)
	findByIDUseCase := adminusecases.NewFindByIDUseCase(&adminRepository)
	findByUserIDUseCase := adminusecases.NewFindByUserIDUseCase(&adminRepository)
	updateUseCase := adminusecases.NewUpdateUseCase(&adminRepository)
	deleteUseCase := adminusecases.NewDeleteUseCase(&adminRepository)

	return services.NewAdminService(
		createUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByUserIDUseCase,
		updateUseCase,
		deleteUseCase,
	)
}

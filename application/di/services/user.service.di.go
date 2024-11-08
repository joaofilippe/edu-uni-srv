package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
	userUseCases "github.com/joaofilippe/edu-uni-srv/domain/usecases/user"
)

func UserServiceFactory(
	userRepository irepositories.IUserRepo,
	adminRepository irepositories.IAdminRepo,
	teacherRepository irepositories.ITeacherRepo,
	studentRepository irepositories.IStudentRepo,
	guardianRepository irepositories.IGuardianRepo,
) *services.UserService {
	createUseCase := userUseCases.NewCreateUseCase(&userRepository)
	loginUseCase := userUseCases.NewLoginUseCase(
		&userRepository,
		&adminRepository,
		&studentRepository,
		&teacherRepository,
		&guardianRepository,
	)
	findAllUseCase := userUseCases.NewFindAllUseCase(&userRepository)
	findByIDUseCase := userUseCases.NewFindByIDUseCase(
		&userRepository,
		&adminRepository,
		&guardianRepository,
		&teacherRepository,
		&studentRepository,
	)
	findByEmailUseCase := userUseCases.NewFindByEmailUseCase(&userRepository)
	updateUseCase := userUseCases.NewUpdateUseCase(&userRepository)
	deleteUseCase := userUseCases.NewDeleteUseCase(&userRepository)

	return services.NewUserService(
		createUseCase,
		loginUseCase,
		findAllUseCase,
		findByIDUseCase,
		findByEmailUseCase,
		updateUseCase,
		deleteUseCase,
	)
}

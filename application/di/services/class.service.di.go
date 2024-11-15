package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
	"github.com/joaofilippe/edu-uni-srv/domain/usecases/class"
)

func ClassServiceFactory(
	classReopsitory irepositories.IClassRepo,
) *services.ClassService {
	createUseCase := classusecases.NewCreateUseCase(classReopsitory)
	return services.NewClassService(createUseCase)
}

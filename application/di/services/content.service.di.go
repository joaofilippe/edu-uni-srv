package servicesdi

import (
	"github.com/joaofilippe/edu-uni-srv/application/services"
	"github.com/joaofilippe/edu-uni-srv/domain/repositories"
	"github.com/joaofilippe/edu-uni-srv/domain/usecases/content"
)

func ContentServiceFactory(
	contentRepository irepositories.IContentRepo,
) *services.ContentService {
	createUseCase := contentusecases.NewCreateUseCase(&contentRepository)

	return services.NewContentService(
		createUseCase,
	)
}

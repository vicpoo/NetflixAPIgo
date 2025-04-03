// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
)

func InitVideoDependencies() (
	*CreateVideoController,
	*GetVideoController,
	*GetAllVideosController, // Nuevo
) {
	repo := NewMySQLVideoRepository()

	createUseCase := application.NewCreateVideoUseCase(repo)
	getByIDUseCase := application.NewGetVideoByIDUseCase(repo)
	getAllUseCase := application.NewGetAllVideosUseCase(repo) // Nuevo

	createController := NewCreateVideoController(createUseCase)
	getController := NewGetVideoController(getByIDUseCase)
	getAllController := NewGetAllVideosController(getAllUseCase) // Nuevo

	return createController, getController, getAllController
}

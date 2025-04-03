package infrastructure

import (
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
)

func InitVideoDependencies() (
	*CreateVideoController,
	*GetVideoController,
) {
	repo := NewMySQLVideoRepository()

	createUseCase := application.NewCreateVideoUseCase(repo)
	getByIDUseCase := application.NewGetVideoByIDUseCase(repo)

	createController := NewCreateVideoController(createUseCase)
	getController := NewGetVideoController(getByIDUseCase)

	return createController, getController
}

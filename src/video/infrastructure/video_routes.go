// video_routes.go
package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
	engine *gin.Engine
}

func NewVideoRouter(engine *gin.Engine) *VideoRouter {
	return &VideoRouter{engine: engine}
}

func (r *VideoRouter) Run() {
	createController, getController, getAllController := InitVideoDependencies() // Actualizado

	videoGroup := r.engine.Group("/videos")
	{
		videoGroup.POST("/", createController.Run)
		videoGroup.GET("/", getAllController.Run) // Nueva ruta
		videoGroup.GET("/:id", getController.Run)
	}
}

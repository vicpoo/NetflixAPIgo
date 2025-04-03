package infrastructure

import (
	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
	engine *gin.Engine
}

func NewVideoRouter(engine *gin.Engine) *VideoRouter {
	return &VideoRouter{
		engine: engine,
	}
}

func (r *VideoRouter) Run() {
	// Inicializar dependencias
	createController, getController := InitVideoDependencies()

	// Configurar rutas
	videoGroup := r.engine.Group("/videos")
	{
		videoGroup.POST("/", createController.Run)
		videoGroup.GET("/:id", getController.Run)
	}
}

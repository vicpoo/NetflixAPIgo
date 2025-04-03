// cache_video_controller.go
package infrastructure

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/video/application"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
)

type CacheVideoController struct {
	cacheService *application.VideoCacheService
	repo         domain.VideoRepository
}

func NewCacheVideoController(
	cacheService *application.VideoCacheService,
	repo domain.VideoRepository,
) *CacheVideoController {
	return &CacheVideoController{
		cacheService: cacheService,
		repo:         repo,
	}
}

func (ctrl *CacheVideoController) CacheVideoHandler(c *gin.Context) {
	// 1. Obtener ID del video
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de video inválido"})
		return
	}

	// 2. Obtener video de la BD
	video, err := ctrl.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}

	// 3. Verificar si ya está cacheados
	if video.IsCacheValid() {
		c.JSON(http.StatusOK, gin.H{
			"message": "El video ya está disponible offline",
			"video":   video,
		})
		return
	}

	// 4. Verificar si es YouTube
	if strings.Contains(video.GetURL(), "youtube.com") || strings.Contains(video.GetURL(), "youtu.be") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No se puede cachear videos de YouTube"})
		return
	}

	// 5. Construir URL completa si es relativa
	videoUrl := video.GetURL()
	if !strings.HasPrefix(videoUrl, "http") {
		videoUrl = "http://localhost:8000" + videoUrl
		video.SetURL(videoUrl)
	}

	// 6. Intentar cachear
	if err := ctrl.cacheService.DownloadVideo(video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Error al descargar el video",
			"details": err.Error(),
		})
		return
	}

	// 7. Actualizar en BD
	if err := ctrl.repo.Save(video); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al actualizar el video",
		})
		return
	}

	// 8. Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Video almacenado para uso offline",
		"video":   video,
	})
}

func (ctrl *CacheVideoController) GetCachedVideoStreamHandler(c *gin.Context) {
	// 1. Obtener ID del video
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de video inválido"})
		return
	}

	// 2. Obtener video de la BD
	video, err := ctrl.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video no encontrado"})
		return
	}

	// 3. Verificar si está cacheados y es válido
	if !video.IsCacheValid() {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "El video no está disponible para visualización offline",
		})
		return
	}

	// 4. Servir el archivo
	c.File(video.GetLocalPath())
}

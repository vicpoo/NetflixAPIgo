// upload_controller.go
package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain"
	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type UploadController struct {
	repo domain.VideoRepository
}

func NewUploadController(repo domain.VideoRepository) *UploadController {
	return &UploadController{repo: repo}
}

func (ctrl *UploadController) UploadHandler(c *gin.Context) {
	// 1. Validar archivo
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Archivo de video requerido"})
		return
	}

	// 2. Validar tipo de archivo
	allowedTypes := map[string]bool{
		".mp4":  true,
		".webm": true,
		".mov":  true,
	}
	ext := filepath.Ext(file.Filename)
	if !allowedTypes[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato no soportado. Use MP4, WebM o MOV"})
		return
	}

	// 3. Crear directorio si no existe
	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear directorio para guardar videos"})
		return
	}

	// 4. Generar nombre único
	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	// 5. Guardar archivo
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el video"})
		return
	}

	// 6. Obtener datos del formulario
	userID, err := strconv.Atoi(c.PostForm("user_id"))
	if err != nil {
		os.Remove(filePath)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
		return
	}

	// 7. Crear registro en BD
	video := &entities.Video{
		Title:       c.PostForm("title"),
		Description: c.PostForm("description"),
		URL:         "/uploads/" + filename, // Ruta accesible públicamente
		UserID:      userID,
	}

	if err := ctrl.repo.Save(video); err != nil {
		os.Remove(filePath)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar en base de datos"})
		return
	}

	// 8. Respuesta exitosa
	c.JSON(http.StatusCreated, gin.H{
		"message": "Video subido exitosamente",
		"video":   video,
		"file":    filename,
	})
}

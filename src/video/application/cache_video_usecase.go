// cache_video_usecase.go
package application

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/vicpoo/NetflixAPIgo/src/video/domain/entities"
)

type VideoCacheService struct {
	CacheDir      string
	CacheDuration time.Duration
}

func NewVideoCacheService(cacheDir string, cacheDuration time.Duration) *VideoCacheService {
	// Crear directorio si no existe
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		panic(fmt.Sprintf("No se pudo crear el directorio de caché: %v", err))
	}

	return &VideoCacheService{
		CacheDir:      cacheDir,
		CacheDuration: cacheDuration,
	}
}

// DownloadVideo descarga y guarda el video localmente para offline
func (s *VideoCacheService) DownloadVideo(video *entities.Video) error {
	// Verificar si es una URL de YouTube
	if isYouTubeURL(video.URL) {
		return fmt.Errorf("no se puede cachear videos de YouTube directamente")
	}

	// Descargar el video
	resp, err := http.Get(video.URL)
	if err != nil {
		return fmt.Errorf("error al descargar el video: %v", err)
	}
	defer resp.Body.Close()

	// Verificar respuesta exitosa
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("respuesta no exitosa: %s", resp.Status)
	}

	// Crear archivo local
	filename := filepath.Join(s.CacheDir, fmt.Sprintf("video_%d%s", video.ID, filepath.Ext(video.URL)))
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error al crear archivo local: %v", err)
	}
	defer file.Close()

	// Copiar contenido
	if _, err := io.Copy(file, resp.Body); err != nil {
		// Si falla, eliminar el archivo parcial
		os.Remove(filename)
		return fmt.Errorf("error al guardar video localmente: %v", err)
	}

	// Actualizar el video
	video.SetLocalPath(filename)
	video.SetIsCached(true)
	video.SetCacheExpiry(time.Now().Add(s.CacheDuration))

	return nil
}

// ClearCache elimina un video del caché
func (s *VideoCacheService) ClearCache(video *entities.Video) error {
	if !video.GetIsCached() || video.GetLocalPath() == "" {
		return nil
	}

	if err := os.Remove(video.GetLocalPath()); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error al eliminar video del caché: %v", err)
	}

	video.SetLocalPath("")
	video.SetIsCached(false)
	video.SetCacheExpiry(time.Time{})

	return nil
}

// isYouTubeURL verifica si la URL es de YouTube
func isYouTubeURL(url string) bool {
	// Implementación simple, puedes mejorarla según tus necesidades
	return strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be")
}

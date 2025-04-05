package main

import (
	"log"
	"os"

	"github.com/vicpoo/NetflixAPIgo/src/core"
	usuarioInfra "github.com/vicpoo/NetflixAPIgo/src/usuario/infrastructure"
	videoInfra "github.com/vicpoo/NetflixAPIgo/src/video/infrastructure"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Inicializar la conexión a la base de datos
	core.InitDB()

	// 2. Crear un router con Gin
	router := gin.Default()

	// 3. Configuración de CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"}, // Asegúrate de permitir la IP pública si lo necesitas
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// 4. Configurar rutas estáticas
	router.Static("/uploads", "./uploads")         // Para videos subidos
	router.Static("/video_cache", "./video_cache") // Para videos cacheados

	// 5. Crear carpetas necesarias
	os.MkdirAll("./uploads", 0755)
	os.MkdirAll("./video_cache", 0755)

	// 6. Middlewares básicos
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 7. Inicializar y configurar rutas
	usuarioRouter := usuarioInfra.NewUsuarioRouter(router)
	usuarioRouter.Run()

	videoRouter := videoInfra.NewVideoRouter(router)
	videoRouter.Run()

	// 8. Iniciar el servidor en todas las interfaces de red
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Printf("\n🚀 Servidor iniciado en http://0.0.0.0:%s", port) // Cambié localhost por 0.0.0.0
	log.Println("📁 Rutas estáticas:")
	log.Println("   - /uploads para videos subidos")
	log.Println("   - /video_cache para videos cacheados")

	// Inicia el servidor en 0.0.0.0:8000 (o el puerto configurado)
	if err := router.Run("0.0.0.0:" + port); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

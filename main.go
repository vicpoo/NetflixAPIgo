package main

import (
	"log"

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

	// 3. Configuración de CORS (permite todas las conexiones)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	// 4. Middlewares básicos
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 5. Inicializar y configurar rutas de usuarios
	usuarioRouter := usuarioInfra.NewUsuarioRouter(router)
	usuarioRouter.Run()

	videoRouter := videoInfra.NewVideoRouter(router)
	videoRouter.Run()

	// 6. Iniciar el servidor en el puerto 8000
	log.Println("API de Usuarios inicializada en http://localhost:8000")
	log.Println("- Rutas de usuarios disponibles en /usuarios")
	log.Println("- Rutas de video dispobibles en /videos")
	log.Println("- Rutas de autenticación disponibles en /auth")

	if err := router.Run(":8000"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuusuario/inmuebles-backend/internal/scraping"
	"github.com/tuusuario/inmuebles-backend/internal/storage"
)

func main() {
	// Inicializamos la base de datos
	storage.InitDB()

	// Inicializamos Gin
	r := gin.Default()

	// Ruta GET para obtener los inmuebles de Argenprop
	r.GET("/inmuebles", func(c *gin.Context) {
		// Hacer scraping de Argenprop
		inmuebles, err := scraping.ScrapeArgenprop()
		if err != nil {
			log.Printf("Error scraping Argenprop: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo inmuebles"})
			return
		}

		// Limitar la respuesta a los primeros 10 inmuebles
		if len(inmuebles) > 10 {
			inmuebles = inmuebles[:10]
		}

		// Devolver los inmuebles en formato JSON
		c.JSON(http.StatusOK, inmuebles)
	})

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Configuración básica del entorno
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
}

// Ejemplo de obtener variables del entorno
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

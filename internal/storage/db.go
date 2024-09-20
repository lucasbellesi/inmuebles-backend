package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Driver para PostgreSQL
)

// Definir la variable db para almacenar la conexión a la base de datos
var db *sql.DB

// Iniciar la conexión a la base de datos
func InitDB() {
	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Obtener la cadena de conexión de la variable de entorno
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL no está definida en las variables de entorno")
	}

	// Abrir la conexión a la base de datos
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error abriendo la conexión a la base de datos: %v", err)
	}

	// Verificar si la conexión es válida
	err = db.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	fmt.Println("Conexión a la base de datos exitosa")
}

// Guardar un inmueble en la base de datos
func SaveInmueble(inmueble Inmueble) error {
	query := `INSERT INTO inmuebles (titulo, precio, ubicacion, link) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, inmueble.Titulo, inmueble.Precio, inmueble.Ubicacion, inmueble.Link)
	if err != nil {
		return err
	}
	return nil
}

// Obtener todos los inmuebles desde la base de datos
func GetAllInmuebles() ([]Inmueble, error) {
	rows, err := db.Query(`SELECT id, titulo, precio, ubicacion, link FROM inmuebles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var inmuebles []Inmueble
	for rows.Next() {
		var inmueble Inmueble
		err := rows.Scan(&inmueble.ID, &inmueble.Titulo, &inmueble.Precio, &inmueble.Ubicacion, &inmueble.Link)
		if err != nil {
			return nil, err
		}
		inmuebles = append(inmuebles, inmueble)
	}

	return inmuebles, nil
}

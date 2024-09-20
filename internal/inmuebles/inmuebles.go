package inmuebles

import (
	"log"
	"strconv"

	"github.com/tuusuario/inmuebles-backend/internal/storage"
)

// Función que recibe los inmuebles scrapeados y los almacena en la base de datos
func ProcessAndStoreInmuebles(inmuebles []storage.Inmueble) error {
	// Validar los datos antes de almacenarlos
	for _, inmueble := range inmuebles {
		if !isValidInmueble(inmueble) {
			log.Printf("Inmueble inválido: %v\n", inmueble)
			continue
		}

		// Almacenar el inmueble en la base de datos
		err := storage.SaveInmueble(inmueble)
		if err != nil {
			return err
		}
	}
	return nil
}

// Validar si los datos del inmueble son correctos
func isValidInmueble(inmueble storage.Inmueble) bool {
	if inmueble.Titulo == "" || inmueble.Precio == "" || inmueble.Link == "" {
		return false
	}
	return true
}

// Obtener todos los inmuebles desde la base de datos
func GetAll() ([]storage.Inmueble, error) {
	inmuebles, err := storage.GetAllInmuebles()
	if err != nil {
		return nil, err
	}
	return inmuebles, nil
}

// Filtrar inmuebles por rango de precios (como ejemplo)
func FilterByPriceRange(minPrice, maxPrice int, inmuebles []storage.Inmueble) []storage.Inmueble {
	var filtered []storage.Inmueble
	for _, inmueble := range inmuebles {
		price := parsePrice(inmueble.Precio) // Función que convierte el precio a un número
		if price >= minPrice && price <= maxPrice {
			filtered = append(filtered, inmueble)
		}
	}
	return filtered
}

// Convierte el precio de string a entero (simple ejemplo)
func parsePrice(priceStr string) int {
	price, _ := strconv.Atoi(priceStr)
	return price
}

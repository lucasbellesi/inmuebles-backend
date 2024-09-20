package storage

// Estructura para representar un inmueble
type Inmueble struct {
	ID        int    `json:"id"`
	Titulo    string `json:"titulo"`
	Precio    string `json:"precio"`
	Ubicacion string `json:"ubicacion"`
	Link      string `json:"link"`
}

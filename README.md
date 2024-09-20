# Proyecto de Scraping de Inmuebles con Golang

Este proyecto realiza scraping de inmuebles en alquiler desde el sitio web [Argenprop](https://www.argenprop.com/departamentos/alquiler) utilizando Golang y la librería [Colly](https://github.com/gocolly/colly). El proyecto expone un único endpoint que devuelve una lista de inmuebles en formato JSON.

## Estructura del Proyecto

inmuebles-backend/ ├── main.go ├── inmuebles/ │ └── argenprop.go ├── go.mod ├── go.sum └── .env

## Requisitos

- Go 1.16 o superior
- PostgreSQL (opcional, si decides almacenar datos)

## Instalación

1. Clona este repositorio:

   ```bash
   git clone https://github.com/tu_usuario/inmuebles-backend.git
   cd inmuebles-backend
   ```

2. Instala las dependencias:

    ```bash
    go mod tidy
    ```

3. Crea un archivo .env en la raíz del proyecto con la siguiente estructura:

    ```bash
    DATABASE_URL=postgres://usuario:contraseña@localhost:5432/nombre_base_datos?sslmode=disable
    ```

## Ejecución

Para ejecutar el proyecto, utiliza el siguiente comando:

```bash
go run main.go
```

El servidor se iniciará en <http://localhost:8080>.

## Endpoint

GET /inmuebles
Este endpoint devuelve una lista de inmuebles en alquiler extraídos de Argenprop. La respuesta será en formato JSON y contendrá los siguientes campos:

titulo: El título del inmueble.
precio: El precio del inmueble.
ubicacion: La ubicación del inmueble.

### Ejemplo de Respuesta

```json
[
    {
        "titulo": "Departamento en Alquiler",
        "precio": "$30,000",
        "ubicacion": "Palermo, Buenos Aires"
    },
    {
        "titulo": "Departamento en Alquiler",
        "precio": "$25,000",
        "ubicacion": "Recoleta, Buenos Aires"
    }
]
```

## Notas

Asegúrate de respetar el archivo robots.txt del sitio web y las políticas de uso.
Si experimentas problemas de acceso o contenido, verifica los selectores en el archivo argenprop.go.

## Contribuciones

Las contribuciones son bienvenidas. Si tienes alguna sugerencia o encuentras un error, no dudes en abrir un issue o enviar un pull request.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.

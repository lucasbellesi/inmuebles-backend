package scraping

import (
	"log"

	"github.com/gocolly/colly"
)

type Inmueble struct {
	Titulo    string
	Precio    string
	Ubicacion string
}

func ScrapeArgenprop() ([]Inmueble, error) {
	var inmuebles []Inmueble

	c := colly.NewCollector()

	c.OnHTML(".card__details-box", func(e *colly.HTMLElement) {
		titulo := e.ChildText(".card__title--primary")
		precio := e.ChildText(".card__price")
		ubicacion := e.ChildText(".card__address")

		inmueble := Inmueble{
			Titulo:    titulo,
			Precio:    precio,
			Ubicacion: ubicacion,
		}

		inmuebles = append(inmuebles, inmueble)
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error durante el scraping:", err)
	})

	err := c.Visit("https://www.argenprop.com/departamentos/alquiler")
	if err != nil {
		return nil, err
	}

	c.Wait()

	return inmuebles, nil
}

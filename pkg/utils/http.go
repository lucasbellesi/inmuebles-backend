package utils

import (
	"net/http"
	"time"
)

// Crear un cliente HTTP con timeout
func CreateHTTPClient() *http.Client {
	return &http.Client{
		Timeout: time.Second * 30,
	}
}

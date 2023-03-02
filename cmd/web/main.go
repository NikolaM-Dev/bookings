package main

import (
	"log"
	"net/http"

	"github.com/NikolaM-Dev/bookings/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}

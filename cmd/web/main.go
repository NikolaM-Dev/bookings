package main

import (
	"log"
	"net/http"

	"github.com/NikolaM-Dev/bookings/pkg/config"
	"github.com/NikolaM-Dev/bookings/pkg/handlers"
	"github.com/NikolaM-Dev/bookings/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting application on port %s\n", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}

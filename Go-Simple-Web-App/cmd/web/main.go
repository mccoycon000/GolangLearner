package main

//This is a simple web app in golang, stuff that is commented out is something I used to test then wanted to remove but still want to show how I did it for reference

import (
	"fmt"
	"log"
	"myGoWebApp/pkg/config"
	"myGoWebApp/pkg/handlers"
	"myGoWebApp/pkg/render"
	"net/http"
)

const portNumber = ":8080"

// main is the main function starts the server on localhost
func main() {
	var app config.AppConfig
	
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Creating template cache failed")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s \n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

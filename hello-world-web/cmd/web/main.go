package main

import (
	"fmt"
	"github.com/lukaszczajkowski/golang-udemy/pkg/config"
	"github.com/lukaszczajkowski/golang-udemy/pkg/handlers"
	"github.com/lukaszczajkowski/golang-udemy/pkg/render"
	"log"
	"net/http"
)

// const cannot be overriden, and var can
const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// start server on port :8080
	// _ = means if there is an error, I don't care
	_ = http.ListenAndServe(portNumber, nil)
}

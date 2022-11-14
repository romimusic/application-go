package main

import (
	"fmt"
	"net/http"

	"github.com/romimusic/application-go/pkg/config"
	"github.com/romimusic/application-go/pkg/handlers"
	"github.com/romimusic/application-go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		fmt.Println("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

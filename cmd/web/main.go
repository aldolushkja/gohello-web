package main

import (
	"fmt"
	"github.com/aldolushkja/gohello-web/config"
	"github.com/aldolushkja/gohello-web/pkg/handlers"
	"github.com/aldolushkja/gohello-web/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

//main this is the main application
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

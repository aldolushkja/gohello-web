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

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

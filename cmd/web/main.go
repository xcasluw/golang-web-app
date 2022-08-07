package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xcasluw/golang-basic-app/pkg/config"
	"github.com/xcasluw/golang-basic-app/pkg/handlers"
	"github.com/xcasluw/golang-basic-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	templaceCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = templaceCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println("Starting application on port", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}

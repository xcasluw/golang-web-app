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

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(":8080", nil)
}

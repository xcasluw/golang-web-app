package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/xcasluw/golang-basic-app/pkg/config"
	"github.com/xcasluw/golang-basic-app/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux
}

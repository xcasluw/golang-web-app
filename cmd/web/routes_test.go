package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/xcasluw/golang-basic-app/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// do nothing. Test passed
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}

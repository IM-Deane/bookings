package main

import (
	"testing"

	"github.com/IM-Deane/bookings/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	router := routes(&app)

	switch v := router.(type) {
	case *chi.Mux:
		// do nothing
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}
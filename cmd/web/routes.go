package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/IM-Deane/bookings/pkg/config"
	"github.com/IM-Deane/bookings/pkg/handlers"
)


func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)


	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)

	return router
}
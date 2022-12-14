package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/IM-Deane/bookings/internal/config"
	"github.com/IM-Deane/bookings/internal/handlers"
)


func routes(app *config.AppConfig) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(NoSurf)
	router.Use(SessionLoad)


	router.Get("/", handlers.Repo.Home)

	router.Get("/about", handlers.Repo.About)

	router.Get("/generals-quarters", handlers.Repo.Generals)

	router.Get("/majors-suite", handlers.Repo.Majors)
	
	router.Get("/search-availability", handlers.Repo.Availability)
	router.Post("/search-availability", handlers.Repo.PostAvailability)
	router.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	router.Get("/contact", handlers.Repo.Contact)


	router.Get("/make-reservation", handlers.Repo.Reservation)
	router.Post("/make-reservation", handlers.Repo.PostReservation)
	router.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	// init static file server
	fileServer := http.FileServer(http.Dir("./static/"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return router
}
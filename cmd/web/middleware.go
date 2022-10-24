package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf enables CSRF protection for POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad load and save session on each request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
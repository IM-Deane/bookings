package handlers

import (
	"net/http"

	"github.com/IM-Deane/bookings/pkg/config"
	"github.com/IM-Deane/bookings/pkg/models"
	"github.com/IM-Deane/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// store users remote IP address in a session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.html", &models.Context{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello friend!"

	// get and store users remote IP from session
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// pass data to template
	render.RenderTemplate(w, "about.page.html", &models.Context{
		StringMap: stringMap,
	})
}
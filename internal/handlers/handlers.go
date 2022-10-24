package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IM-Deane/bookings/internal/config"
	"github.com/IM-Deane/bookings/internal/forms"
	"github.com/IM-Deane/bookings/internal/models"
	"github.com/IM-Deane/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.page.html", &models.Context{})
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
	render.RenderTemplate(w, r, "about.page.html", &models.Context{
		StringMap: stringMap,
	})
}


// Reservation is the Reservation page and displays form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.Context{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handles the posting of the reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	
	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName: r.Form.Get("last_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}

	// init form scheme object
	form := forms.New(r.PostForm)

	// validate submitted fields
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		// re-render page with errors
		render.RenderTemplate(w, r, "make-reservation.page.html", &models.Context{
			Form: form,
			Data: data,
		})
		return
	}
}


// Generals is the general's quarters page handler
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.Context{
	})
}

// Majors is the major's suite page handler
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.Context{
	})
}

// Availability is the make reservation page handler
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.Context{
	})
}

// PostAvailability is the make reservation page handler
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	// get form data
	startDate := r.Form.Get("start-date")
	endDate := r.Form.Get("end-date")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", startDate, endDate)))
}

type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handles request for availability and returns JSON response
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	res := jsonResponse{
		OK: true,
		Message: "Avaliable!",
	}

	// create JSON with fields from struct
	out, err := json.MarshalIndent(res, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact displays the website's contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.Context{
	})
}
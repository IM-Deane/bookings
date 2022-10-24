package models

import "github.com/IM-Deane/bookings/internal/forms"

// Context holds data passed from handlers to templates
type Context struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{} // use this when unsure of other data types
	CSRFToken string
	Flash string // success
	Warning string
	Error string
	Form *forms.Form
}
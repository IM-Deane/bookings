package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/IM-Deane/bookings/pkg/config"
	"github.com/IM-Deane/bookings/pkg/models"
)


var app *config.AppConfig

// NewTemplates sets the config for the template pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData returns template data
func AddDefaultData(context *models.Context) *models.Context {
	// TODO: update this with additional data
	return context
}

// RenderTemplate renders page using HTML template
func RenderTemplate(w http.ResponseWriter, tmpl string, context *models.Context) {

	var tempCache map[string]*template.Template
	if app.UseCache {
		// read info from cache
		tempCache = app.TemplateCache
	} else {
		// rebuild cache
		tempCache, _ = CreateTemplateCache()
	}
	
	t, ok := tempCache[tmpl]
	if !ok {
		log.Fatal("couldn't get template from cache")
	}

	buf := new(bytes.Buffer)

	context = AddDefaultData(context)

	err := t.Execute(buf, context)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

// createTemplateCache using templates from "templates/*"
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all template files named "*.page.html" BEFORE loading template layout
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// iterate through all files ending with *.page.html
	for _, page := range pages {
		// get filename
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page) // ie. home.page.html
		if err != nil {
			return myCache, err
		}

		// now we get the layout files
		matches, err := filepath.Glob("./templates/*layout.html")
		if err != nil {
			return myCache, err
		}
		
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
			return myCache, err
		}
		}

		// cache template
		myCache[name] = ts
	}

	return myCache, nil
}

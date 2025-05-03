package ui

import (
	"html/template"
	"log"
	"net/http"

	"github.com/furkankarayel/URL_Shortener/internal/api"
	"github.com/furkankarayel/URL_Shortener/internal/utils"
)

type UI struct {
	templates *template.Template
}

func New() *api.Route {
	UIHandler := &UI{}
	UIHandler.templates = template.Must(template.ParseGlob("./templates/*.html"))
	return &api.Route{
		WithLogger: true,
		Handler:    UIHandler,
	}
}

func (s *UI) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)

	switch head {
	case "":
		s.serveTemplate(w, r, "index.html")
	default:
		api.Respond(w, r, http.StatusNotFound, "Oh uh.. what you're looking for is not here..")

	}
}

func (s *UI) serveTemplate(w http.ResponseWriter, r *http.Request, templateName string) {
	// Execute the template
	err := s.templates.ExecuteTemplate(w, templateName, nil)
	if err != nil {
		log.Printf("Error executing template %s: %v", templateName, err)
		api.Respond(w, r, http.StatusInternalServerError, "Server error")
	}
}

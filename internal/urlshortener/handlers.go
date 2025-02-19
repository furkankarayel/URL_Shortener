package urlshortener

import (
	"html/template"
	"net/http"

	"github.com/furkankarayel/URL_Shortener/internal/api"
)

func (s *URLService) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		api.Respond(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return

	}

	parseErr := r.ParseForm()
	if parseErr != nil {
		api.Respond(w, r, http.StatusBadRequest, "Failed to parse form data")
		return
	}

	url := r.FormValue("url")
	if url == "" {
		api.Respond(w, r, http.StatusBadRequest, "URL is required")
		return
	}

	shortenedURL, err := s.shortenURL(url)
	if err != nil {
		api.Respond(w, r, http.StatusInternalServerError, "Failed to shorten URL"+err.Error())
		return
	}

	tmpl := template.Must(template.New("result").Parse(`
	<div id="result">
		<p>Your shortened URL is: <a href="{{.ShortURL}}">{{.ShortURL}}</a></p>
	</div>
	`))

	data := struct {
		ShortURL string
	}{
		ShortURL: shortenedURL,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
	}

}

func (s *URLService) GetOriginalURL(w http.ResponseWriter, r *http.Request, shortCode string) {
	if r.Method != http.MethodGet {
		api.Respond(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return

	}

	originalURL, err := s.getLongURL(shortCode)
	if err != nil {
		api.Respond(w, r, http.StatusInternalServerError, "Failed to get original URL"+err.Error())
		return
	}

	api.Redirect(w, r, originalURL, http.StatusMovedPermanently)

}

package urlshortener

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/furkankarayel/URL_Shortener/internal/api"
	"github.com/furkankarayel/URL_Shortener/internal/utils"
)

func (s *URLService) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {

		api.Respond(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return

	}

	var req struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		api.Respond(w, r, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.URL == "" {
		api.Respond(w, r, http.StatusBadRequest, "URL is required")
		return
	}

	shortURL, err := s.shortenURL(req.URL)
	if err != nil {
		api.Respond(w, r, http.StatusInternalServerError, "Failed to shorten URL"+err.Error())
		return
	}

	api.Respond(w, r, http.StatusOK, shortURL)

}

func (s *URLService) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		api.Respond(w, r, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var shortCode string
	shortCode, r.URL.Path = utils.ShiftPath(r.URL.Path)
	log.Println(shortCode)

	originalURL, err := s.getLongURL(shortCode)
	if err != nil {
		api.Respond(w, r, http.StatusInternalServerError, "Failed to get original URL"+err.Error())
		return
	}

	api.Respond(w, r, http.StatusOK, originalURL)

}

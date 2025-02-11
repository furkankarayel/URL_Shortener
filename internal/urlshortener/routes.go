package urlshortener

import (
	"database/sql"
	"net/http"

	"github.com/furkankarayel/URL_Shortener/internal/api"
	"github.com/furkankarayel/URL_Shortener/internal/cache"
	"github.com/furkankarayel/URL_Shortener/internal/utils"
)

type URLHandler struct {
	URLService *URLService
}

func New(db *sql.DB, urlCache *cache.URLCache) *api.Route {
	URLHandler := &URLHandler{
		URLService: NewURLService(db, urlCache),
	}
	return &api.Route{
		WithLogger: true,
		Handler:    URLHandler,
	}
}

func (s *URLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)

	switch head {
	case "":
		api.Respond(w, r, http.StatusOK, "Nothing here")
	case "shorten":
		s.URLService.CreateShortURL(w, r)
	default:
		s.URLService.GetOriginalURL(w, r, head)

	}
}

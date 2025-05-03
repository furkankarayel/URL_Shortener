package api

import (
	"encoding/json"
	"net/http"

	"github.com/furkankarayel/URL_Shortener/internal/api/middleware"
	"github.com/furkankarayel/URL_Shortener/internal/utils"
)

type Route struct {
	WithLogger bool
	Handler    http.Handler
}

type Server struct {
	Routes map[string]*Route
}

func New(routes map[string]*Route) *Server {
	return &Server{Routes: routes}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = utils.ShiftPath(r.URL.Path)
	route, ok := s.Routes[head]
	if !ok {
		Respond(w, r, http.StatusNotFound, "root route not found")
		return
	}

	next := route.Handler

	if route.WithLogger {
		next = middleware.Logger(next)
	}
	next.ServeHTTP(w, r)
}

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {
	if e, ok := data.(error); ok {
		var tmp = new(struct {
			Status string `json:"status"`
			Error  string `json:"error"`
		})
		tmp.Status = "error"
		tmp.Error = e.Error()
		data = tmp
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	middleware.LogRequest(r, status)

	return nil
}

func Redirect(w http.ResponseWriter, r *http.Request, url string, status int) {
	http.Redirect(w, r, url, status)
	middleware.LogRequest(r, status)
}

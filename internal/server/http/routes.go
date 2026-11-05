package httpserver

import (
	"net/http"
)

func (s *Server) addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", s.health)
}

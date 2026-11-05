package httpserver

import (
	"net/http"
)

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("rapidcache http server is running!"))
	s.logger.Debug("GET /health called")
}

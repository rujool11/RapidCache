package httpserver

import (
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", health)
}

package httpserver

import (
	"log/slog"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("rapidcache http server is running!"))
	slog.Debug("GET /health called")
}

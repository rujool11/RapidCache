package httpserver

import "net/http"

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("rapidcache http server is running!"))
}

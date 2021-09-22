package httputil

import (
	"io"
	"net/http"
)

func WriteString(w http.ResponseWriter, status int, s string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	io.WriteString(w, s)
}

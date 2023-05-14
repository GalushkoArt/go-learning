package utils

import (
	"log"
	"net/http"
)

type loggingHandler struct {
	next http.Handler
}

func NewLoggingHandler(next http.Handler) http.Handler {
	return &loggingHandler{next: next}
}

func (h *loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s\n", r.Method, r.URL)
	h.next.ServeHTTP(w, r)
}

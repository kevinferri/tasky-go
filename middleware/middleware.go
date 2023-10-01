package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitMiddleware(router *mux.Router) {
	router.Use(loggingMiddleware, contentTypeApplicationJsonMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Host)

		next.ServeHTTP(w, r)
	})
}

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
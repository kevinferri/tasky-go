package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitHandlers(r *mux.Router) {
	// health_check
	r.HandleFunc("/health_check", getHealthCheck).Methods(http.MethodGet)

	// snippets
	r.HandleFunc("/snippets", getSnippet).Methods(http.MethodGet)
	r.HandleFunc("/snippets/new", postSnippet).Methods(http.MethodPost)

	log.Println("✅ Handlers")
}

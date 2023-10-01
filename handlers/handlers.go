package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitHandlers(r *mux.Router) {
	// health_check
	r.HandleFunc("/health_check", getHealthCheck).Methods(http.MethodGet)

	// notes
	r.HandleFunc("/notes", getNote).Methods(http.MethodGet)
	r.HandleFunc("/notes/new", postNote).Methods(http.MethodPost)

}

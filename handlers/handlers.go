package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func InitHandlers(r *mux.Router) {
	// health_check
	r.HandleFunc("/health_check", getHealthCheck).Methods(http.MethodGet)

	// snippets
	r.HandleFunc("/snippets", getAllSnippets).Methods(http.MethodGet)
	r.HandleFunc("/snippets/{id}", getSnippetById).Methods(http.MethodGet)
	r.HandleFunc("/snippets/new", postSnippet).Methods(http.MethodPost)

	log.Println("✅ Handlers")
}

func handleErrResp(i interface{}, w http.ResponseWriter, err error) {
	t := reflect.TypeOf(i)

	if err.Error() == sql.ErrNoRows.Error() {
		http.Error(w, t.Name()+" not found", http.StatusNotFound)

	}

	log.Println("Handler error:", err.Error())
	http.Error(w, "Bad request", http.StatusBadRequest)
}

package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/kevinferri/tasky-go/store"
)

type Handler struct {
	router       *mux.Router
	snippetStore *store.SnippetStore
}

func NewHandler(router *mux.Router, db *sqlx.DB) *Handler {
	return &Handler{
		router:       router,
		snippetStore: store.NewSnippetStore(db),
	}
}

func InitHandlers(h *Handler) {
	// health_check
	h.router.HandleFunc("/health_check", h.getHealthCheck).Methods(http.MethodGet)

	// snippets
	h.router.HandleFunc("/snippets", h.getAllSnippets).Methods(http.MethodGet)
	h.router.HandleFunc("/snippets", h.postSnippet).Methods(http.MethodPost)
	h.router.HandleFunc("/snippets/{id}", h.getSnippetById).Methods(http.MethodGet)

	log.Println("✅ Handlers")
}

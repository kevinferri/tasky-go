package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinferri/tasky-go/response"
	"github.com/kevinferri/tasky-go/store"
)

func (h *Handler) getAllSnippets(w http.ResponseWriter, r *http.Request) {
	snippets, err := h.snippetStore.All()

	if err != nil {
		response.HandleErr(w, store.Snippet{}, err)
	}

	response.JSON(w, snippets, http.StatusOK)
}

func (h *Handler) getSnippetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	s, err := h.snippetStore.ById(params["id"])

	if err != nil {
		response.HandleErr(w, s, err)
		return
	}

	response.JSON(w, s, http.StatusOK)
}

func (h *Handler) postSnippet(w http.ResponseWriter, r *http.Request) {
	s := &store.Snippet{}
	err := json.NewDecoder(r.Body).Decode(s)

	if err != nil {
		response.HandleErr(w, s, err)
	}

	h.snippetStore.Create(s)
	response.JSON(w, s, http.StatusCreated)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinferri/tasky-go/store"
)

func getAllSnippets(w http.ResponseWriter, r *http.Request) {
	var s store.Snippet
	snippets, err := store.GetAllSnippets()

	if err != nil {
		handleErrResp(s, w, err)
	}

	json.NewEncoder(w).Encode(snippets)
}

func getSnippetById(w http.ResponseWriter, r *http.Request) {
	var s store.Snippet
	params := mux.Vars(r)
	s, err := store.GetSnippetById(params["id"])

	if err != nil {
		handleErrResp(s, w, err)
		return
	}

	json.NewEncoder(w).Encode(s)
}

func postSnippet(w http.ResponseWriter, r *http.Request) {
	var s store.Snippet
	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		handleErrResp(s, w, err)
	}

	store.CreateSnippet(s)
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/kevinferri/tasky-go/store"
)

func getSnippet(w http.ResponseWriter, r *http.Request) {
	var s store.Snippet

	snippets, _ := s.GetAll()

	json.NewEncoder(w).Encode(snippets)
}

func postSnippet(w http.ResponseWriter, r *http.Request) {
	var s store.Snippet

	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Create()
}

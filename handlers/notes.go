package handlers

import (
	"net/http"

	"github.com/kevinferri/tasky-go/types"
)

func getNote(w http.ResponseWriter, r *http.Request) {
	resp := &types.Note{
		ID:    "980123908123",
		Title: "Hello, world!",
	}

	w.Write(resp.ToJson())
}

func postNote(w http.ResponseWriter, r *http.Request) {
	// todo
}

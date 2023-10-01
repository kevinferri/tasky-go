package handlers

import (
	"encoding/json"
	"net/http"
)

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

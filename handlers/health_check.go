package handlers

import (
	"net/http"

	"github.com/kevinferri/tasky-go/response"
)

func getHealthCheck(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, map[string]string{"status": "ok"}, http.StatusOK)
}

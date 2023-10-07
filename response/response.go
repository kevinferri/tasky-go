package response

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"reflect"
)

type Response struct {
	Data     interface{} `json:"data,omitempty"`
	ErrorMsg string      `json:"error,omitempty"`
	Status   int         `json:"status,omitempty"`
}

func JSON(w http.ResponseWriter, v interface{}, s int) error {
	r := &Response{
		Data:   v,
		Status: s,
	}

	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(r)
}

func HandleErr(w http.ResponseWriter, v interface{}, err error) {
	t := reflect.TypeOf(v)
	r := &Response{}

	switch err.Error() {
	case sql.ErrNoRows.Error():
		r.ErrorMsg = t.Name() + " not found"
		r.Status = http.StatusNotFound
	default:
		r.ErrorMsg = "Bad request"
		r.Status = http.StatusBadRequest
		log.Println("Unhandled handler error:", err.Error())
	}

	w.WriteHeader(r.Status)
	json.NewEncoder(w).Encode(r)
}

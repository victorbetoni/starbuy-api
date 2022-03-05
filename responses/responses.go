package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

func Error(w http.ResponseWriter, status int, err error) {
	JSON(w, status, struct {
		Error string `json:"error"`
	}{Error: err.Error()})
}

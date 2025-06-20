package handler

import (

	"encoding/json"
	"net/http"
	
)

type Health struct {
	Message string `json:"message"`
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Health{"ok"})
	
}
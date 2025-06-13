package router

import (
	"github.com/gorilla/mux"
	handler "github.com/paulo-fabiano/simple-crud-api/internal/handler/taskHandler"
)

func loadRoutesAPI(router *mux.Router) {

	// Routes API
	router.HandleFunc("/tasks", handler.ListTasks).Methods("GET") 
	router.HandleFunc("/task", nil).Methods("GET")
	router.HandleFunc("/task", nil).Methods("POST")
	router.HandleFunc("/task/{id}", nil).Methods("PUT", "PATCH", "DELETE")

}
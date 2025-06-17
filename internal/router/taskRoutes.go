package router

import (
	
	"github.com/gorilla/mux"
	handler "github.com/paulo-fabiano/simple-crud-api/internal/handler/taskHandler"
	health "github.com/paulo-fabiano/simple-crud-api/internal/handler/health"

)

func loadRoutesAPI(router *mux.Router) {

	// Routes API
	router.HandleFunc("/health", health.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/tasks", handler.ListTasksHandler).Methods("GET") 
	router.HandleFunc("/task/{id}", handler.ListTaskHandler).Methods("GET")
	router.HandleFunc("/task", handler.CreateTask).Methods("POST")
	router.HandleFunc("/task/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/task/{id}", handler.UpdateTaskHandler).Methods("PATCH")

}
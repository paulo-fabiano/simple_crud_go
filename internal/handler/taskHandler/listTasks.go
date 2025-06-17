package handler

import (

	"encoding/json"
	"net/http"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"

)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks := repository.ListTasks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	w.WriteHeader(http.StatusOK)

}
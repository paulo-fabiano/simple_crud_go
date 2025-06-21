package handler

import (
	"encoding/json"
	"net/http"

	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
)

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := repository.ListTasks()
	if err != nil {
		sendError(w, http.StatusInternalServerError, "erro interno")
		return 
	}

	sendListTasks(w, http.StatusOK, tasks)

}

type SendListTasks struct {
	Data	interface{}	`json:"data"`
}

type SendTask struct {
	Data	interface{}	`json:"data"`
}

func sendListTasks(writer http.ResponseWriter, code int, data interface{}) {

	writer.Header().Set("Content-Type", "aplication/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(SendListTasks{
		Data: data,
	})

}

func sendListTask(writer http.ResponseWriter, code int, data interface{}) {

	writer.Header().Set("Content-Type", "aplication/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(SendTask{
		Data: data,
	})

}




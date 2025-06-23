package handler

import (
	
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



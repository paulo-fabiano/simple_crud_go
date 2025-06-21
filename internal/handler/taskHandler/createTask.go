package handler

import (

	"encoding/json"
	"net/http"

	"github.com/paulo-fabiano/simple-crud-api/internal/model"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
	
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var taskRequest model.TaskRequest
	taskDecoder := json.NewDecoder(r.Body)
	err := taskDecoder.Decode(&taskRequest)
	if err != nil {
		logger.Errorf("Erro ao fazer o decode: &v", err)
		sendError(w, http.StatusInternalServerError, "error in decode object")
		return
	}
	
	err = repository.CreateTask(&taskRequest)
	if err != nil {
		logger.Errorf("Erro ao criar objeto no banco de dados: 5v", err)
		sendError(w, http.StatusInternalServerError, "error creating object")
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create task!"))	
	
}
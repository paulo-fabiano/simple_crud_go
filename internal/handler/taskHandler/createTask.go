package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/paulo-fabiano/simple-crud-api/internal/model"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var taskRequest model.TaskRequest
	taskDecoder := json.NewDecoder(r.Body)
	err := taskDecoder.Decode(&taskRequest)
	if err != nil {
		log.Fatal("Error in decode", err)
	}
	
	// Save
	err = repository.CreateTask(&taskRequest)
	if err != nil {
		log.Fatal(err.Error())
	}
	// w.Header().Set("Content-Type", "application/josn")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Create task!"))	
	
}
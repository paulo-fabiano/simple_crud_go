package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/simple-crud-api/internal/model"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idTask, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	var task model.TaskRequest

	taskDecoder := json.NewDecoder(r.Body)
	err = taskDecoder.Decode(&task)
	if err != nil {
		log.Fatal(err)
	}

	
	err = repository.UpdateTask(idTask, task)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error update task"))
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task Update"))

}
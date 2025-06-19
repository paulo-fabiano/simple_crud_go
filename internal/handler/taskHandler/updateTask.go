package handler

import (

	"encoding/json"
	"log"
	"net/http"

	"github.com/paulo-fabiano/simple-crud-api/internal/model"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"

)

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	idString, err := getIDParams(r)
	if err != nil {
		sendError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	id, err := convertIDToInt(*idString)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	var task model.TaskRequest

	taskDecoder := json.NewDecoder(r.Body)
	err = taskDecoder.Decode(&task)
	if err != nil {
		log.Fatal(err)
	}
	
	err = repository.UpdateTask(*id, task)
	if err != nil {
		sendError(w, http.StatusInternalServerError, "erro ao atualizar o objeto")
		return
	}

	sendSucess(w, http.StatusOK, "task atualizada")

}
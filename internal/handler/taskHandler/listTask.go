package handler

import (
	"net/http"

	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
)

func ListTaskHandler(w http.ResponseWriter, r *http.Request) {

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
	
	taskResponse, err := repository.ListTask(*id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendListTask(w, http.StatusOK, taskResponse)

}
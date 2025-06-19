package handler

import (
	"net/http"

	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

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
	
	err = repository.DeleteTask(id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(w, http.StatusOK, "objeto deletado com sucesso")

}
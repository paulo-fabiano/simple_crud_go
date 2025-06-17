package handler

import (

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"

)

func ListTaskHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	idTask, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("erro id não é um número"))
		return
	}
	taskResponse := repository.ListTask(idTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(taskResponse)
	w.WriteHeader(http.StatusOK)

}
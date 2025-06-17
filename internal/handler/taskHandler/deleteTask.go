package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	repository "github.com/paulo-fabiano/simple-crud-api/internal/repository/task"
)

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]
	idTask, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	repository.DeleteTask(&idTask)

	// implementar resto da função
}
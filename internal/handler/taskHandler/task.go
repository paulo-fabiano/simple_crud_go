package handler

import (

	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/paulo-fabiano/simple-crud-api/internal/config"
	
)

var (
	logger *config.Logger
)

type MessageError struct {
	Message 	string		`json:"message"`
	DateTime 	time.Time	`json:"date"`
}

func sendError(w http.ResponseWriter, code int, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(MessageError{
		Message:  message,
		DateTime: time.Now(),
	})
	
}

type MessageSucess struct {
	Message		string		`json:"message"`	
}

func sendSucess(w http.ResponseWriter, code int, message string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(MessageSucess{
		Message: message,
	})

}

func getIDParams(request *http.Request) (*string, error) {

	parametros := mux.Vars(request)
	iD := parametros["id"]

	if iD == "" {
		return nil, fmt.Errorf("error ID is empty")
	}
	
	return &iD, nil

}

func convertIDToInt(idString string) (*int, error) {

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, fmt.Errorf("error in convert ID to type int")
	}

	return &id, nil

}

type SendTask struct {
	Data	interface{}	`json:"data"`
}

func sendListTask(writer http.ResponseWriter, code int, data interface{}) {

	writer.Header().Set("Content-Type", "aplication/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(SendTask{
		Data: data,
	})

}

type SendListTasks struct {
	Data	interface{}	`json:"data"`
}

func sendListTasks(writer http.ResponseWriter, code int, data interface{}) {

	writer.Header().Set("Content-Type", "aplication/json")
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(SendListTasks{
		Data: data,
	})

}

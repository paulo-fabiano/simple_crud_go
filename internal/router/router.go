package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitializeRoutesAPI() {

	portServer := os.Getenv("PORT_SERVER_API")
	if portServer == "" {
		portServer = "8080"
	}

	router := mux.NewRouter()

	// Inicializando task routes
	loadRoutesAPI(router)

	err := http.ListenAndServe(":"+portServer, router)
	if err != nil {
		log.Fatal("error creating http server")
	}

}
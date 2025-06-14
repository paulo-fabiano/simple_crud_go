package main

import (
	
	"log"

	"github.com/paulo-fabiano/simple-crud-api/internal/config"
	"github.com/paulo-fabiano/simple-crud-api/internal/router"
)

func main() {

	// Inicializando DB
	err := config.SetupDB()
	if err != nil {
		log.Fatal("Error in setup config databse: ", err)
	}
	
	log.Printf("Aplicação rodando na porta 8080")

	// Inicializando Server e Routes
	router.InitializeRoutesAPI()


}
package main

// Alteração Local
// Comentário Remoto

import (
	
	"log"

	"github.com/paulo-fabiano/simple-crud-api/internal/config"
	"github.com/paulo-fabiano/simple-crud-api/internal/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Inicializando DB
	err := config.SetupDB()
	if err != nil {
		logger.Errorf("Error in setup config databse: %v", err)
		panic(err)
	}
	logger.Info("Banco de dados inicializado com sucesso")	
	log.Printf("Aplicação rodando na porta 8080")

	// Inicializando Server e Routes
	router.InitializeRoutesAPI()

}

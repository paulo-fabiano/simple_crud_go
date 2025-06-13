package config

import (

	"fmt"
	"log"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

)

var (
	dbConnection *sql.DB
)

// Inicializa a conexão com o banco de dados
func SetupDB() (error) {

	// Carrega as variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		
		dbHost string = os.Getenv("DB_HOST")
		dbPort string = os.Getenv("DB_PORT")
		dbUser string = os.Getenv("DB_USERNAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName string = os.Getenv("DB_DATABASE")
		
	)
	
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	dbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error opening connection to database", err)
		return err
	}
	
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal("Error in ping command on database", err)
		return err
	}

	return nil

}

// Retorna a conexão para ser usada em outros locais
func GetConnectionDB() *sql.DB {
	return dbConnection
}
package config

import (

	"fmt"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

)

var (
	DBConnection *sql.DB
	logger *Logger
)

// Inicializa a conexão com o banco de dados
func SetupDB() (error) {

	logger := GetLogger("SetupDB")

	// Carrega as variáveis do arquivo .env
	err := godotenv.Load()
	if err != nil {
		logger.Error("Error loading .env file")
	}

	var (
		
		dbHost string = os.Getenv("DB_HOST")
		dbPort string = os.Getenv("DB_PORT")
		dbUser string = os.Getenv("DB_USERNAME")
		dbPassword string = os.Getenv("DB_PASSWORD")
		dbName string = os.Getenv("DB_DATABASE")
		
	)
	
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	DBConnection, err = sql.Open("postgres", connectionString)
	if err != nil {
		logger.Errorf("Error opening connection to database: %v", err)
		return err
	}
	
	err = DBConnection.Ping()
	if err != nil {
		logger.Errorf("Error in ping command on database: %v", err)
		return err
	}

	return nil

}

// Retorna a conexão para ser usada em outros locais
func GetConnectionDB() *sql.DB {
	return DBConnection
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
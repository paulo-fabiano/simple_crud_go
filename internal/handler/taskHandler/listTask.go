package handler

import (

	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/paulo-fabiano/simple-crud-api/internal/config"
	"github.com/paulo-fabiano/simple-crud-api/internal/model"

)

func ListTask(w http.ResponseWriter, r *http.Request, id *uint) {

	// Buscando a conexão com o DB
	db := config.GetConnectionDB()

	query := "SELECT * FROM tasks WHERE id = $;"
	row := db.QueryRow(query, &id)

	// Variável que irá receber os valores
	var task model.TaskResponse
	err := row.Scan(&task.ID, &task.CreatedAt, &task.DeletedAt, &task.UpdatedAt, 
		&task.Name, &task.Description, &task.Status)
	defer row.Scan()
	
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error id not exists"))
			return 
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error in query"))
			return 
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
	w.WriteHeader(http.StatusOK)

}
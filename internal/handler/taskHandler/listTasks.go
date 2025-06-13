package handler

import (

	"encoding/json"
	"net/http"

	"github.com/paulo-fabiano/simple-crud-api/internal/config"
	"github.com/paulo-fabiano/simple-crud-api/internal/model"

)

func ListTasks(w http.ResponseWriter, r *http.Request) {

	// Buscando a conexão com o DB
	db := config.GetConnectionDB()
	rows, err := db.Query("SELECT * FROM tasks;")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return 
	}
	defer rows.Close() // Find why

	var tasks []model.TaskResponse

	for rows.Next() {

		var task model.TaskResponse
		err := rows.Scan(&task.ID, &task.CreatedAt, &task.DeletedAt, &task.UpdatedAt ,&task.Name, &task.Description, &task.Status)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return 
		}

		tasks = append(tasks, task)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	w.WriteHeader(http.StatusOK)

}
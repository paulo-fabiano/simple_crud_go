package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/paulo-fabiano/simple-crud-api/internal/config"
	model "github.com/paulo-fabiano/simple-crud-api/internal/model"
)

var (
	logger *config.Logger
)

func ListTask(id int) (*model.TaskResponse, error) {

	dbConn := config.GetConnectionDB()
	stmt, err := dbConn.Prepare("SELECT * FROM tasks WHERE id = $1")
	if err != nil {
		logger.Errorf("Erro no prepare statement: &v", err)
	}
	defer stmt.Close()

	var taskResponse model.TaskResponse
	err = stmt.QueryRow(id).Scan(&taskResponse.ID, &taskResponse.CreatedAt,
	&taskResponse.UpdatedAt, &taskResponse.DeletedAt, &taskResponse.Name, 
	&taskResponse.Description, &taskResponse.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ID não existe no database")
		} else {
			return nil, fmt.Errorf("erro desconhecido")
		}
	}

	return &taskResponse, nil

}

func ListTasks() ([]model.TaskResponse, error) {

	db := config.GetConnectionDB()
	rows, err := db.Query("SELECT * FROM tasks;")
	if err != nil {
		logger.Errorf("Erro ao buscar tasks no banco de dados: %v", err)
		return nil, fmt.Errorf("erro no banco de dados")
	}
	defer rows.Close() 

	var tasks []model.TaskResponse

	for rows.Next() {

		var task model.TaskResponse
		err := rows.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt, &task.DeletedAt, &task.Name, &task.Description, &task.Status)
		if err != nil {
			logger.Errorf("Erro ao fazer Scan no objeto: %v", err)
			return nil, fmt.Errorf("Erro ao fazer Scan em um dos objetos da lista")
		}

		tasks = append(tasks, task)

	}

	return tasks, nil
}

func CreateTask(task *model.TaskRequest) error {

	dbConn := config.GetConnectionDB()
	stmt, err := dbConn.Prepare("INSERT INTO tasks (name, created_at, description, status) VALUES ($1, $2, $3, $4);")
	if err != nil {
		logger.Errorf("Erro ao criar task no banco de dados: %v", err)
		return fmt.Errorf("Erro ao criar objeto no banco de dados")
	}
	defer stmt.Close()

	_ = stmt.QueryRow(task.Name, time.Now(), task.Description, task.Status)
	if err != nil {
		logger.Errorf("Error ao fazer a query no banco de dados: %v", err)
		return fmt.Errorf("Erro ao criar objeto no banco de dados")
	}

	return nil

}

func UpdateTask(id *int, task model.TaskRequest)  (*model.TaskResponse, error) {

	db := config.GetConnectionDB()
	stmt, err := db.Prepare("SELECT EXISTS(SELECT id FROM tasks WHERE id = $1)")
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao buscar dados no banco de dados")
	}
	defer stmt.Close()

	var existe bool
	err = stmt.QueryRow(id).Scan(&existe)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao buscar dados no banco de dados")
	}

	if !existe {
		return nil, fmt.Errorf("ID não existe no banco de dados")
	}

	// Se existe
	stmt, err = db.Prepare("SELECT * FROM tasks WHERE id = $1;")
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao buscar dados do usuário")
	}

	var taskE model.Task
	
	err = stmt.QueryRow(id).Scan(&taskE.ID, &taskE.CreatedAt, &taskE.UpdatedAt, &taskE.DeletedAt, &taskE.Name, &taskE.Description, &taskE.Status)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao buscar dados do usuário")
	}

	if task.Name != "" {
		taskE.Name = task.Name
	}

	if task.Description != "" {
		taskE.Description = task.Description
	}

	if task.Status == false {
		taskE.Status = taskE.Status
	}

	stmt, err = db.Prepare("UPDATE tasks SET created_at = $1, updated_at = $2, deleted_at = $3, name = $4, description = $5, status = $6 WHERE id = $7")
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao atualizar dados do objeto")
	}

 	_, err = stmt.Exec(taskE.CreatedAt, time.Now(), taskE.DeletedAt, taskE.Name, taskE.Description, taskE.Status, id)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao atualizar dados do objeto")
	}

	stmt, err = db.Prepare("SELECT id, created_at, updated_at, deleted_at, name, description, status FROM tasks WHERE id = $1")
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao buscar o objeto atualizado")
	}

	var taskUpdated *model.TaskResponse
	err = stmt.QueryRow(id).Scan(&taskUpdated.ID, &taskUpdated.CreatedAt, &taskUpdated.UpdatedAt, &taskUpdated.DeletedAt, &taskUpdated.Name, &taskUpdated.Description, &taskUpdated.Status)
	if err != nil {
		logger.Error(err)
		return nil, fmt.Errorf("Erro ao buscar o objeto atualizado")
	}


	return taskUpdated, nil
}

func DeleteTask(id *int) error {

	db := config.GetConnectionDB()
	stmt, err := db.Prepare("SELECT EXISTS(SELECT id FROM tasks WHERE id = $1)")
	if err != nil {
		return fmt.Errorf("Error in prepare statement")
	}
	defer stmt.Close()

	var existe bool
	err = stmt.QueryRow(id).Scan(&existe)
	if err != nil {
		return fmt.Errorf("Error in query database")
	}

	if !existe {
		return fmt.Errorf("Error ID not exist")		
	}

	stmtDelete, err := db.Prepare("DELETE FROM tasks WHERE id = $1")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Error in prepare statement")
	}
	defer stmtDelete.Close()
	_, err = stmtDelete.Exec(id)
	if err != nil {
		return fmt.Errorf("Error in query database")
	}

	return nil

}
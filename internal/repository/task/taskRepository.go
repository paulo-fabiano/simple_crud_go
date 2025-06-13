package repository

import (

	"github.com/paulo-fabiano/simple-crud-api/internal/model"

)

type Repository interface {

	ListTasks(id uint) []model.TaskResponse
	ListTask(id uint) model.TaskResponse
	CreateTask(task model.Task) model.TaskResponse

	// No momento deixar só esses métodos mesmo, meia noite fi kkk. Dormir
}

// Implementar funções abaixo
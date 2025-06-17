package model

import (
	"database/sql"
	"time"
)

type Task struct {
	ID uint	`json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt sql.NullTime `json:"updatedAt"`
	DeletedAt sql.NullTime `json:"deteledAt"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status bool `json:"status"`
}

//Dto Request
type TaskRequest struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Status bool `json:"status"`
}

// Dto Response
type TaskResponse struct {
	ID uint	`json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deteledAt"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status bool `json:"status"`
}
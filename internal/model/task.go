package model

import "time"

type Task struct {
	ID uint	`json:"id"`
	Name string `json:"nome"`
	Description string `json:"description"`
	Status bool `json:"status"`
}

// Dto Response
type TaskResponse struct {
	ID uint	`json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Name string `json:"nome"`
	Description string `json:"description"`
	Status bool `json:"status"`
}
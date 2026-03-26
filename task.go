package main

import "time"

type TaskStatus string

const (
	StatusPending    TaskStatus = "Pendiente"
	StatusInProgress TaskStatus = "En Curso"
	StatusDone       TaskStatus = "Hecho"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

package main

import (
	"fmt"
	"time"
)

const msgNotFound = "no se encontró tarea con ID %d"

// AddTask agrega una nueva tarea y retorna la lista actualizada junto con la tarea creada.
func AddTask(tasks []Task, description string) ([]Task, Task, error) {
	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	newTask := Task{
		ID:          newID,
		Description: description,
		Status:      "Pendiente",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	tasks = append(tasks, newTask)
	return tasks, newTask, nil
}

// DeleteTask elimina la tarea con el ID dado y retorna la lista actualizada.
func DeleteTask(tasks []Task, id int) ([]Task, error) {
	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, fmt.Errorf(msgNotFound, id)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return tasks, nil
}

// UpdateTask actualiza la descripción de la tarea con el ID dado.
func UpdateTask(tasks []Task, id int, newDesc string) ([]Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return nil, fmt.Errorf(msgNotFound, id)
}

// MarkStatus cambia el estado de la tarea con el ID dado.
func MarkStatus(tasks []Task, id int, status string) ([]Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			return tasks, nil
		}
	}
	return nil, fmt.Errorf(msgNotFound, id)
}

package main

import (
	"fmt"
	"time"
)

const msgNotFound = "no se encontró tarea con ID %d"

type TaskManager struct {
	userID int
	tasks  []Task
}

func NewTaskManager(userID int) (*TaskManager, error) {
	tasks, err := LoadTasks(userID)
	if err != nil {
		return nil, err
	}
	return &TaskManager{userID: userID, tasks: tasks}, nil
}

// save wraps the SaveTasks call
func (m *TaskManager) save() error {
	return SaveTasks(m.userID, m.tasks)
}

func (m *TaskManager) AddTask(description string) (Task, error) {
	newID := 1
	for _, t := range m.tasks {
		if t.ID >= newID {
			newID = t.ID + 1
		}
	}

	newTask := Task{
		ID:          newID,
		UserID:      m.userID,
		Description: description,
		Status:      StatusPending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	m.tasks = append(m.tasks, newTask)
	return newTask, m.save()
}

func (m *TaskManager) DeleteTask(id int) error {
	index := -1
	for i, t := range m.tasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf(msgNotFound, id)
	}

	m.tasks = append(m.tasks[:index], m.tasks[index+1:]...)
	return m.save()
}

func (m *TaskManager) UpdateTask(id int, newDesc string) error {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks[i].Description = newDesc
			m.tasks[i].UpdatedAt = time.Now()
			return m.save()
		}
	}
	return fmt.Errorf(msgNotFound, id)
}

func (m *TaskManager) MarkStatus(id int, status TaskStatus) error {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks[i].Status = status
			m.tasks[i].UpdatedAt = time.Now()
			return m.save()
		}
	}
	return fmt.Errorf(msgNotFound, id)
}

func (m *TaskManager) GetTasks() []Task {
	return m.tasks
}

package main

import (
	"fmt"
	"strconv"
)

const (
	msgInvalidID = "ID inválido"
	msgError     = "Error:"
)

func handleAdd(args []string, manager *TaskManager) {
	if len(args) < 1 {
		fmt.Println("Error: Debes proporcionar una descripción. Ejemplo: taskGo add \"Mi tarea\"")
		return
	}

	newTask, err := manager.AddTask(args[0])
	if err != nil {
		fmt.Println("Error al crear tarea:", err)
		return
	}

	fmt.Printf("Tarea agregada exitosamente (ID %d)\n", newTask.ID)
}

func handleList(args []string, manager *TaskManager) {
	filter := ""
	if len(args) > 0 {
		filter = args[0]
	}

	fmt.Printf("%-4s | %-12s | %s\n", "ID", "Estado", "Descripción")
	fmt.Println("--------------------------------------------------")
	for _, t := range manager.GetTasks() {
		if filter == "" || filter == string(t.Status) {
			fmt.Printf("%-4d | %-12s | %s\n", t.ID, t.Status, t.Description)
		}
	}
}

func handleUpdate(args []string, manager *TaskManager) {
	if len(args) < 2 {
		fmt.Println("Uso: update <id> <nueva descripción>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(msgInvalidID)
		return
	}

	err = manager.UpdateTask(id, args[1])
	if err != nil {
		fmt.Println(msgError, err)
		return
	}

	fmt.Printf("Tarea %d actualizada\n", id)
}

func handleDelete(args []string, manager *TaskManager) {
	if len(args) < 1 {
		fmt.Println("Uso: delete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(msgInvalidID)
		return
	}

	err = manager.DeleteTask(id)
	if err != nil {
		fmt.Println(msgError, err)
		return
	}

	fmt.Printf("Tarea %d eliminada\n", id)
}

func handleMarkStatus(args []string, manager *TaskManager, newStatus TaskStatus) {
	if len(args) < 1 {
		fmt.Printf("Uso: mark <id> (para marcar como %s)\n", newStatus)
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(msgInvalidID)
		return
	}

	err = manager.MarkStatus(id, newStatus)
	if err != nil {
		fmt.Println(msgError, err)
		return
	}

	fmt.Printf("Tarea %d marcada como %s\n", id, newStatus)
}

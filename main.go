package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Error: Se requiere un comando. Ejemplos: add, list, update, delete")
		os.Exit(1)
	}

	command := args[0]

	cmdArgs := args[1:]

	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("Error al cargar tareas:", err)
		os.Exit(1)
	}

	switch command {
	case "add":
		handleAdd(cmdArgs, tasks)
	case "list":
		handleList(cmdArgs, tasks)
	case "update":
		handleUpdate(cmdArgs, tasks)
	case "delete":
		handleDelete(cmdArgs, tasks)
	case "mark-in-progress":
		handleMarkStatus(cmdArgs, tasks, "En Curso")
	case "mark-done":
		handleMarkStatus(cmdArgs, tasks, "Hecho")
	default:
		fmt.Printf("Comando desconocido: %s\n", command)
	}
}

func handleAdd(args []string, tasks []Task) {
	if len(args) < 1 {
		fmt.Println("Error: Debes proporcionar una descripción. Ejemplo: taskGo add \"Mi tarea\"")
		return
	}

	description := args[0]

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

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error al guardar la tarea:", err)
		return
	}
	fmt.Printf("Tarea agregada exitosamente (ID %d)\n", newTask.ID)
}

func handleList(args []string, tasks []Task) {
	filter := ""
	if len(args) > 0 {

		filter = args[0]
	}

	fmt.Printf("%-4s | %-12s | %s\n", "ID", "Estado", "Descripción")
	fmt.Println("--------------------------------------------------")
	for _, t := range tasks {
		if filter == "" || filter == t.Status {
			fmt.Printf("%-4d | %-12s | %s\n", t.ID, t.Status, t.Description)
		}
	}
}

func handleUpdate(args []string, tasks []Task) {
	if len(args) < 2 {
		fmt.Println("Uso: update <id> <nueva descripción>")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ID inválido")
		return
	}
	newDesc := args[1]

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = newDesc
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("No se encontró tarea con ID %d\n", id)
		return
	}

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error al guardar:", err)
		return
	}
	fmt.Printf("Tarea %d actualizada\n", id)
}

func handleDelete(args []string, tasks []Task) {
	if len(args) < 1 {
		fmt.Println("Uso: delete <id>")
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	index := -1
	for i, t := range tasks {
		if t.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Printf("No se encontró tarea con ID %d\n", id)
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error al guardar:", err)
		return
	}
	fmt.Printf("Tarea %d eliminada\n", id)
}

func handleMarkStatus(args []string, tasks []Task, newStatus string) {
	if len(args) < 1 {
		fmt.Printf("Uso: %s <id>\n", newStatus)
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("ID inválido")
		return
	}

	found := false
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("No se encontró tarea con ID %d\n", id)
		return
	}

	if err := SaveTasks(tasks); err != nil {
		fmt.Println("Error al guardar:", err)
		return
	}
	fmt.Printf("Tarea %d marcada como %s\n", id, newStatus)
}

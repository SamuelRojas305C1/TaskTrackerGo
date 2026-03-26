package main

import (
	"fmt"
	"os"
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


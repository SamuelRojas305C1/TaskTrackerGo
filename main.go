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

	manager, err := NewTaskManager()
	if err != nil {
		fmt.Println("Error al iniciar TaskManager:", err)
		os.Exit(1)
	}

	switch command {
	case "add":
		handleAdd(cmdArgs, manager)
	case "list":
		handleList(cmdArgs, manager)
	case "update":
		handleUpdate(cmdArgs, manager)
	case "delete":
		handleDelete(cmdArgs, manager)
	case "mark-in-progress":
		handleMarkStatus(cmdArgs, manager, StatusInProgress)
	case "mark-done":
		handleMarkStatus(cmdArgs, manager, StatusDone)
	default:
		fmt.Printf("Comando desconocido: %s\n", command)
	}
}

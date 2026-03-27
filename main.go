package main

import (
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Uso: tasktrackergo <comando> [argumentos]")
	fmt.Println("Comandos de Autenticación:")
	fmt.Println("  register <email> <password> <name>  Registra un nuevo usuario")
	fmt.Println("  login <email> <password>            Inicia sesión")
	fmt.Println("  logout                              Cierra sesión")
	fmt.Println("Comandos de Tareas (requieren inicio de sesión):")
	fmt.Println("  add <descripcion>    Agrega una nueva tarea")
	fmt.Println("  list [estado]        Lista las tareas (opcional: filtra por estado)")
	fmt.Println("  update <id> <desc>   Actualiza la descripción de una tarea")
	fmt.Println("  delete <id>          Elimina una tarea")
	fmt.Println("  mark-in-progress <id> Marca una tarea como 'En Curso'")
	fmt.Println("  mark-done <id>       Marca una tarea como 'Hecho'")
	fmt.Println("\nPara obtener ayuda sobre un comando, usa: tasktrackergo <comando> --help")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]
	if command == "help" || command == "--help" || command == "-h" {
		printUsage()
		os.Exit(0)
	}

	cmdArgs := os.Args[2:]

	switch command {
	case "register":
		handleRegister(cmdArgs)
		return
	case "login":
		handleLogin(cmdArgs)
		return
	case "logout":
		handleLogout()
		return
	}

	session, err := GetCurrentSession()
	if err != nil {
		fmt.Println("Error: debes iniciar sesión para usar este comando. Usa 'login' primero.")
		os.Exit(1)
	}

	manager, err := NewTaskManager(session.UserID)
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
		printUsage()
		os.Exit(1)
	}
}

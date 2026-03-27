package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const (
	msgInvalidID = "ID inválido"
	msgError     = "Error:"
)

func handleRegister(args []string) {
	cmd := flag.NewFlagSet("register", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Println("Uso: tasktrackergo register <email> <password> <name>")
	}
	cmd.Parse(args)

	if cmd.NArg() < 3 {
		cmd.Usage()
		return
	}

	email := cmd.Arg(0)
	password := cmd.Arg(1)
	name := strings.Join(cmd.Args()[2:], " ")

	user, err := RegisterUser(email, password, name)
	if err != nil {
		fmt.Println("Error al registrar usuario:", err)
		return
	}

	fmt.Printf("Usuario registrado exitosamente. Bienvenido/a, %s! (ID %d)\n", user.Name, user.ID)
	// We don't auto login, let them login.
}

func handleLogin(args []string) {
	cmd := flag.NewFlagSet("login", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Println("Uso: tasktrackergo login <email> <password>")
	}
	cmd.Parse(args)

	if cmd.NArg() < 2 {
		cmd.Usage()
		return
	}

	email := cmd.Arg(0)
	password := cmd.Arg(1)

	session, err := LoginUser(email, password)
	if err != nil {
		fmt.Println("Error al iniciar sesión:", err)
		return
	}

	fmt.Printf("Inicio de sesión exitoso. Tu ID de sesión es %d\n", session.UserID)
}

func handleLogout() {
	err := LogoutUser()
	if err != nil {
		fmt.Println("Error al cerrar sesión:", err)
		return
	}
	fmt.Println("Sesión cerrada exitosamente.")
}

func handleAdd(args []string, manager *TaskManager) {
	cmd := flag.NewFlagSet("add", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Println("Uso: tasktrackergo add <descripción>")
	}
	cmd.Parse(args)

	if cmd.NArg() < 1 {
		cmd.Usage()
		return
	}

	description := strings.Join(cmd.Args(), " ")
	newTask, err := manager.AddTask(description)
	if err != nil {
		fmt.Println("Error al crear tarea:", err)
		return
	}

	fmt.Printf("Tarea agregada exitosamente (ID %d)\n", newTask.ID)
}

func handleList(args []string, manager *TaskManager) {
	cmd := flag.NewFlagSet("list", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Println("Uso: tasktrackergo list [Pendiente|En Curso|Hecho]")
	}
	cmd.Parse(args)

	filter := cmd.Arg(0)

	fmt.Printf("%-4s | %-12s | %s\n", "ID", "Estado", "Descripción")
	fmt.Println("--------------------------------------------------")
	for _, t := range manager.GetTasks() {
		if filter == "" || filter == string(t.Status) {
			fmt.Printf("%-4d | %-12s | %s\n", t.ID, t.Status, t.Description)
		}
	}
}

func handleUpdate(args []string, manager *TaskManager) {
	cmd := flag.NewFlagSet("update", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Println("Uso: tasktrackergo update <id> <nueva descripción>")
	}
	cmd.Parse(args)

	if cmd.NArg() < 2 {
		cmd.Usage()
		return
	}

	id, err := strconv.Atoi(cmd.Arg(0))
	if err != nil {
		fmt.Println(msgInvalidID)
		return
	}

	newDesc := strings.Join(cmd.Args()[1:], " ")
	err = manager.UpdateTask(id, newDesc)
	if err != nil {
		fmt.Println(msgError, err)
		return
	}

	fmt.Printf("Tarea %d actualizada\n", id)
}

func handleDelete(args []string, manager *TaskManager) {
	cmd := flag.NewFlagSet("delete", flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Println("Uso: tasktrackergo delete <id>")
	}
	cmd.Parse(args)

	if cmd.NArg() < 1 {
		cmd.Usage()
		return
	}

	id, err := strconv.Atoi(cmd.Arg(0))
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
	cmdName := "mark-in-progress"
	if newStatus == StatusDone {
		cmdName = "mark-done"
	}
	cmd := flag.NewFlagSet(cmdName, flag.ExitOnError)
	cmd.Usage = func() {
		fmt.Printf("Uso: tasktrackergo %s <id>\n", cmdName)
	}
	cmd.Parse(args)

	if cmd.NArg() < 1 {
		cmd.Usage()
		return
	}

	id, err := strconv.Atoi(cmd.Arg(0))
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

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	usersFile   = "users.json"
	sessionFile = "session.json"
	dataDir     = "data"
)

func getTasksFileName(userID int) string {
	return filepath.Join(dataDir, fmt.Sprintf("tasks_%d.json", userID))
}

func ensureDataDir() error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		return os.MkdirAll(dataDir, 0755)
	}
	return nil
}

func LoadTasks(userID int) ([]Task, error) {
	err := ensureDataDir()
	if err != nil {
		return nil, err
	}

	fileName := getTasksFileName(userID)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(userID int, tasks []Task) error {
	err := ensureDataDir()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}

	fileName := getTasksFileName(userID)
	return os.WriteFile(fileName, data, 0644)
}

func LoadUsers() ([]User, error) {
	if _, err := os.Stat(usersFile); os.IsNotExist(err) {
		return []User{}, nil
	}

	data, err := os.ReadFile(usersFile)
	if err != nil {
		return nil, err
	}

	var users []User
	err = json.Unmarshal(data, &users)
	return users, err
}

func SaveUsers(users []User) error {
	data, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(usersFile, data, 0644)
}

func LoadSession() (*Session, error) {
	if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("no hay sesión activa")
	}

	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return nil, err
	}

	var session Session
	err = json.Unmarshal(data, &session)
	return &session, err
}

func SaveSession(session Session) error {
	data, err := json.MarshalIndent(session, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(sessionFile, data, 0644)
}

func ClearSession() error {
	if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(sessionFile)
}

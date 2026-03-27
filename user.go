package main

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Name         string `json:"name"`
}

type Session struct {
	UserID int `json:"userId"`
}

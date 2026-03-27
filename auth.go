package main

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(email, password, name string) (User, error) {
	users, err := LoadUsers()
	if err != nil {
		return User{}, err
	}

	for _, u := range users {
		if u.Email == email {
			return User{}, errors.New("el correo ya está registrado")
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	newID := 1
	for _, u := range users {
		if u.ID >= newID {
			newID = u.ID + 1
		}
	}

	newUser := User{
		ID:           newID,
		Email:        email,
		PasswordHash: string(hashedPassword),
		Name:         name,
	}

	users = append(users, newUser)
	err = SaveUsers(users)
	return newUser, err
}

func LoginUser(email, password string) (Session, error) {
	users, err := LoadUsers()
	if err != nil {
		return Session{}, err
	}

	var foundUser *User
	for _, u := range users {
		if u.Email == email {
			foundUser = &u
			break
		}
	}

	if foundUser == nil {
		return Session{}, errors.New("credenciales inválidas")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password))
	if err != nil {
		return Session{}, errors.New("credenciales inválidas")
	}

	session := Session{UserID: foundUser.ID}
	err = SaveSession(session)
	return session, err
}

func LogoutUser() error {
	return ClearSession()
}

func GetCurrentSession() (*Session, error) {
	return LoadSession()
}

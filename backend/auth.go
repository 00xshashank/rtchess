package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), 10)
}

func checkPasswordHash(password string, hash string) (bool, error) {
	pass_hash, pass_err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if pass_err != nil {
		return false, pass_err
	}

	return (string(pass_hash) == hash), nil
}

func signup(user User) error {
	hashed_password, pwd_err := hashPassword(user.Password)
	if pwd_err != nil {
		return pwd_err
	}

	user.Password = string(hashed_password)
	return AddUser(user)
}

func checkUserValidity(user User) error {
	if len(user.Username) < 6 {
		return errors.New("Username length should be at least 6")
	}

	if len(user.Password) < 8 {
		return errors.New("Password length must be at least 8")
	}
	if user.Firstname == "" {
		return errors.New("Invalid first name")
	}

	if user.Lastname == "" {
		return errors.New("Inalid last name")
	}

	if user.Email == "" {
		return errors.New("Invalid email")
	}

	return nil
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	body := make([]byte, 0)
	p := make([]byte, 8)
	for {
		clear(p)
		n, err := r.Body.Read(p)
		if err != nil && err.Error() != "EOF" {
			log.Printf("Error while reading request body: %s\n", err)
			w.WriteHeader(500)
			return
		}
		body = append(body, p[:n]...)
		if n < 8 {
			break
		}
	}

	var user User
	json_err := json.Unmarshal(body, &user)
	if json_err != nil {
		log.Printf("Error while decoding JSON in signup handler: %s\n", json_err)
		w.WriteHeader(400)
		return
	}

	log.Printf("Decoded user: %v", user)

	check_err := checkUserValidity(user)
	if check_err != nil {
		w.WriteHeader(400)
		w.Write([]byte(check_err.Error()))
		return
	}

	signup_err := signup(user)
	if signup_err != nil {
		fmt.Printf("Error while signng up user: %s\n", signup_err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
}

type loginRequest struct {
	username string
	password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}

	body := make([]byte, 0)
	p := make([]byte, 8)
	for {
		clear(p)
		n, err := r.Body.Read(p)
		if err != nil && err.Error() != "EOF" {
			log.Printf("Error while reading request body: %s\n", err)
			w.WriteHeader(500)
			return
		}
		body = append(body, p[:n]...)
		if n < 8 {
			break
		}
	}

	var user loginRequest
	json_err := json.Unmarshal(body, &user)
	if json_err != nil {
		log.Printf("Error while decoding JSON in login handler: %s\n", json_err)
		w.WriteHeader(400)
		return
	}

	// pass_match, pass_err := checkPasswordHash(user.password, )
}

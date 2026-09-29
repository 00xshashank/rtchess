package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/00xshashank/rtchess/backend/db"
	"github.com/00xshashank/rtchess/backend/types"
)

func signup(user types.User) error {
	hashed_password, pwd_err := hashPassword(user.Password)
	if pwd_err != nil {
		return pwd_err
	}

	user.Password = string(hashed_password)
	return db.AddUser(user)
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

	var user types.User
	json_err := json.Unmarshal(body, &user)
	if json_err != nil {
		log.Printf("Error while decoding JSON in signup handler: %s\n", json_err)
		w.WriteHeader(400)
		return
	}

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

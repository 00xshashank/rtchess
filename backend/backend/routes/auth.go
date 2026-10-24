package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/00xshashank/rtchess/backend/backend/auth"
	"github.com/00xshashank/rtchess/backend/backend/db"
	"github.com/00xshashank/rtchess/backend/backend/types"
)

var AuthMuxer http.ServeMux

func signup(user types.UserModel) error {
	hashed_password, pwd_err := auth.HashPassword(user.Password)
	if pwd_err != nil {
		return pwd_err
	}

	user.Password = string(hashed_password)
	return db.AddUserToDB(user)
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

	var user types.UserModel
	json_err := json.Unmarshal(body, &user)
	if json_err != nil {
		log.Printf("Error while decoding JSON in signup handler: %s\n", json_err)
		w.WriteHeader(400)
		return
	}

	check_err := auth.CheckUserValidity(user)
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
	Username string
	Password string
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

	var login loginRequest
	json_err := json.Unmarshal(body, &login)
	if json_err != nil {
		log.Printf("Error while decoding JSON in login handler: %s\n", json_err)
		w.WriteHeader(400)
		return
	}

	user, user_err := db.GetUserFromDB(login.Username)
	if user_err != nil {
		log.Printf("Error while retrieving user for login: %s\n", user_err)
		w.WriteHeader(200)
		w.Write([]byte("false"))
		return
	}

	pass_match := auth.CheckPasswordHash([]byte(login.Password), []byte(user.Password))

	if pass_match {
		token, token_err := auth.CreateJWT(user)
		if token_err != nil {
			log.Printf("Error while getting JWT: %s\n", token_err)

			return
		}

		cookie := http.Cookie{
			Name:  "chesslocal",
			Value: token,
		}

		http.SetCookie(w, &cookie)

		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	cookie, cookie_err := r.Cookie("chesslocal")
	if cookie_err != nil {
		log.Printf("Error while fetching cookie from request: %s\n", cookie_err)
		w.Write([]byte(cookie_err.Error()))
		w.WriteHeader(403)
	}

	user, user_err := auth.GetUserFromCookie(cookie)
	if user_err != nil {
		log.Printf("Error while getting user: %s\n", user_err)
		w.Write([]byte(user_err.Error()))
		w.WriteHeader(403)
	}

	json_user, json_err := json.Marshal(user)
	if json_err != nil {
		log.Printf("Error while encoding JSON: %s\n", json_err)
		w.Write([]byte(json_err.Error()))
		w.WriteHeader(403)
	}

	w.Write([]byte(json_user))
}

func init() {
	AuthMuxer.HandleFunc("/signup", SignupHandler)
	AuthMuxer.HandleFunc("/login", LoginHandler)
	AuthMuxer.HandleFunc("/getuser", GetUserHandler)
}

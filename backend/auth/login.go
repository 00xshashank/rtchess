package auth

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/00xshashank/rtchess/backend/db"
)

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

	pass_match := checkPasswordHash([]byte(login.Password), []byte(user.Password))

	if pass_match {
		token, token_err := getJWT(user)
		if token_err != nil {
			log.Printf("Error while getting JWT: %s\n", token_err)

			return
		}

		cookie := http.Cookie{
			Name:  "chesslocal",
			Value: token,
		}

		http.SetCookie(w, &cookie)
		log.Printf("Set cookie: %s\n", &cookie)

		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}
}

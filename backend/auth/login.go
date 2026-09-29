package auth

import (
	"encoding/json"
	"log"
	"net/http"
)

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

package routes

import (
	"log"
	"net/http"

	"github.com/00xshashank/rtchess/backend/auth"
)

var AuthMuxer http.ServeMux

func init() {
	AuthMuxer.HandleFunc("/signup", auth.SignupHandler)
	AuthMuxer.HandleFunc("/login", auth.LoginHandler)
	AuthMuxer.HandleFunc("/getuser", auth.GetUserHandler)

	AuthMuxer.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	log.Printf("Registered all routes for AuthMuxer.\n")
}

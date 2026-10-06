package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/00xshashank/rtchess/backend/types"
	"github.com/golang-jwt/jwt/v5"
)

func GetUserFromCookie(token *jwt.Token) (types.User, error) {
	exp_time, exp_err := token.Claims.GetExpirationTime()
	if exp_err != nil {
		log.Printf("Error while getting expiration time: %s\n", exp_err)
		return types.User{}, exp_err
	}

	if exp_time.Time.Unix() < time.Now().Unix() {
		return types.User{}, errors.New("Session Expired")
	}

	var user types.User = types.User{}

	return user, nil
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	cookie, cookie_err := r.Cookie("chesslocal")
	if cookie_err != nil {
		log.Printf("Error while fetching cookie from request: %s\n", cookie_err)
		w.Write([]byte(cookie_err.Error()))
		w.WriteHeader(403)
	}

	token, token_err := jwt.ParseWithClaims(cookie.Value, &jwtClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if token_err != nil {
		log.Printf("Error while decoding token: %s\n", token_err)
		w.Write([]byte(token_err.Error()))
		w.WriteHeader(403)
	}

	user, user_err := GetUserFromCookie(token)
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

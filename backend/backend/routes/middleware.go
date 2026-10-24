package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/00xshashank/rtchess/backend/backend/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Executing from auth middleware\n")

			log.Printf("%v\n", *r)

			cookie, cookie_err := r.Cookie("chesslocal")
			if cookie_err != nil {
				log.Printf("Error while getting cookie from request: %s\n", cookie_err)
				w.WriteHeader(403)
				return
			}

			user, user_err := auth.GetUserFromCookie(cookie)
			if user_err != nil {
				log.Printf("Error while getting user from cookie: %s\n", user_err)
				w.WriteHeader(403)
				return
			}

			log.Printf("No user errors present\n")

			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)

			log.Printf("Done executing auth middleware\n")

			next.ServeHTTP(w, r)
		},
	)
}

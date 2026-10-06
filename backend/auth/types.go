package auth

import "github.com/golang-jwt/jwt/v5"

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Username  string `json:"usr"`
	Email     string `json:"eml"`
	Firstname string `json:"fnm"`
	Lastname  string `json:"lnm"`
}

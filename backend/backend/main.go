package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/00xshashank/rtchess/backend/backend/db"
	"github.com/00xshashank/rtchess/backend/backend/routes"
	"github.com/00xshashank/rtchess/backend/backend/ws"
	"github.com/joho/godotenv"
)

func init() {
	env_err := godotenv.Load()
	if env_err != nil {
		fmt.Printf("Error while loading environment variables: %s\n", env_err)
		return
	}

	db_err := db.InitDB()
	if db_err != nil {
		fmt.Printf("Failed to initialize database: %s\n", db_err)
		os.Exit(124)
	}

	db.InitRedis()
}

func main() {
	PORT := 3000

	muxer := http.NewServeMux()
	muxer.Handle("/auth/{wildcard}", http.StripPrefix("/auth", &routes.AuthMuxer))
	muxer.Handle("/game/{wildcard}", http.StripPrefix("/game", &routes.GameMuxer))
	muxer.Handle("/join", routes.AuthMiddleware(http.HandlerFunc(ws.JoinGameHandler)))

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: muxer,
	}

	log.Printf("Starting server on port: %d\n", PORT)
	log.Fatal(server.ListenAndServe())
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	env_err := godotenv.Load()
	if env_err != nil {
		fmt.Printf("Error while loading environment variables: %s\n", env_err)
		return
	}

	db_err := InitDB()
	if db_err != nil {
		fmt.Printf("Failed to initialize database: %s\n", db_err)
		os.Exit(124)
	}
}

func main() {
	PORT := 3000

	muxer := http.NewServeMux()
	muxer.HandleFunc("/signup", SignupHandler)
	muxer.HandleFunc("/login", LoginHandler)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: muxer,
	}

	fmt.Printf("Starting server on port: %d\n", PORT)
	log.Fatal(server.ListenAndServe())
}

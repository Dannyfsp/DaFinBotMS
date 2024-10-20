package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Welcome to my server")
	})

	mux.HandleFunc("/{$}", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "404 Not Found")
	})

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

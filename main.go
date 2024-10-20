package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dannyfsp/DaFinBotMS/database"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

}

func main() {
	connUri := os.Getenv("MONGO_URI")

	ctx, err := database.ConnectDB(connUri)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Defer the disconnection from MongoDB
	defer func() {
		if err := database.Client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Server is Up and running")
	})

	mux.HandleFunc("/{$}", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "404 Not Found")
	})

	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

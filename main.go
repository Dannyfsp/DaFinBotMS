package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Dannyfsp/DaFinBotMS/database"
	"github.com/Dannyfsp/DaFinBotMS/middleware"
	"github.com/Dannyfsp/DaFinBotMS/routes"
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

	if err := database.CreateCollections(); err != nil {
		log.Fatalf("Failed to create collections: %v", err)
	}

	router := http.NewServeMux()

	routes.LoadRoutes(router)

	stack := middleware.CreateStack(
		middleware.LoggingMiddleware,
		middleware.CorsMiddleware,
	)

	port := os.Getenv("PORT")

	server := http.Server{
		Addr:    port,
		Handler: stack(router),
	}

	fmt.Println("Server is running on port", port)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Error starting server:", err)
	}
}

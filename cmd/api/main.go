package main

import (
	"log"
	// "net/http"
	"github.com/Vladmir-dev/fintech-wallet/internal/db"
	"github.com/Vladmir-dev/fintech-wallet/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file (if it exists)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found → using system environment variables (normal in Docker)")
	}

	// Connect to database
	db.Connect()

	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router, db.DB)

	// Start server
	log.Println("🚀 Server running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

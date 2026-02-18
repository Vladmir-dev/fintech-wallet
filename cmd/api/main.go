package main

import (
	"log"
	// "net/http"
	"github.com/Vladmir-dev/fintech-wallet/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/Vladmir-dev/fintech-wallet/internal/routes"
)

func main() {
	// Create Gin router
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
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

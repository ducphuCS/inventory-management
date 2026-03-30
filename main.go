package main

import (
	"log"

	"github.com/ducphu/inventory-service/handlers"
	"github.com/ducphu/inventory-service/storage"
	"github.com/gin-gonic/gin"
)

// main is the application entry point.
// It initializes the database storage, setup HTTP routes, and starts the server.
func main() {
	// 1. Initialize SQLite Database
	dbPath := "inventory.db"
	dbStorage, err := storage.NewSQLiteStorage(dbPath)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}
	defer dbStorage.Close() // Ensure DB is closed when the application exits

	// 2. Initialize HTTP Handlers
	inventoryHandler := handlers.NewItemHandler(dbStorage)

	// 3. Initialize the Gin framework and HTTP Router
	router := gin.Default()

	// 4. Define API versioned routes
	v1Routes := router.Group("/v1")
	{
		// Endpoints for inventory management
		v1Routes.POST("/items", inventoryHandler.AddItem)
		v1Routes.GET("/items/count", inventoryHandler.GetTotal)
		v1Routes.GET("/items/:id/analysis", inventoryHandler.GetAnalysis)
	}

	// 5. Setup basic health check for the service
	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{"status": "ok"})
	})

	// 6. Start the service
	log.Println("Starting inventory-service on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

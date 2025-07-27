package main

import (
	"log"

	"gocommerce/pkg/config"

	"gocommerce/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load config and DB
	config.LoadEnv()
	db.ConnectPostgres()

	// Initialize Gin
	r := gin.Default()

	// Health check route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start the server
	log.Printf("Server starting on port %s...", config.Cfg.ServerPort)
	err := r.Run(":" + config.Cfg.ServerPort)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

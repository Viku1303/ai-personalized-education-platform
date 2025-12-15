package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/ai-personalized-education-platform-go/api"
	"github.com/yourusername/ai-personalized-education-platform-go/database"
)

func main() {
	// Initialize database connection
	err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	r := gin.Default()

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "AI Personalized Education Platform is running!",
		})
	})

	// Authentication routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", api.Register)
		auth.POST("/login", api.Login)
	}

	// Welcome route
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to AI-Powered Education Platform API",
			"version": "1.0",
			"endpoints": gin.H{
				"auth":   "/api/auth/register, /api/auth/login",
				"status": "Server is running successfully",
			},
		})
	})

	// Start server on port 8080
	log.Println("Server starting on port 8080...")
	// Get port from environment (Railway sets PORT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s...", port)
	r.Run(":" + port)
}

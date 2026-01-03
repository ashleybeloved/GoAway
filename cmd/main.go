package main

import (
	"goaway/internal/handlers"
	"goaway/internal/repositories"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env

	err := godotenv.Load("configs/.env")
	if err != nil {
		log.Fatal("Failed to load .env file:", err)
	}

	// Connect to PostgreSQL Database

	err = repositories.StartPostgreSQL()
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	log.Println("Connecting to PostgreSQL database successfully")

	// Set Release Mode for Gin

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Handlers

	r.POST("/reg", handlers.Reg)
	r.POST("/login", handlers.Login)
	r.POST("/logout")

	r.POST("/new")
	r.GET("/:id")

	// Run server

	log.Println("The server is running on :8080")
	r.Run(":" + os.Getenv("PORT"))
}

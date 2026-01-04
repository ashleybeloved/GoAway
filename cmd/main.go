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

	// Connect to PostgreSQL

	err = repositories.StartPostgreSQL()
	if err != nil {
		log.Fatal("Could not connect to PostgreSQL:", err)
	}
	log.Println("Connecting to PostgreSQL successfully")

	// Connect to Redis

	err = repositories.StartRedis()
	if err != nil {
		log.Fatal("Could not connect to Redis:", err)
	}
	log.Println("Connecting to Redis successfully")

	// Run server

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/:id", handlers.Redirect)
	r.POST("/reg", handlers.Reg)
	r.POST("/login", handlers.Login)
	r.POST("/logout", handlers.Logout)

	r.POST("/new", handlers.New)

	log.Println("The server is running on :8080")
	r.Run(":" + os.Getenv("PORT"))
}

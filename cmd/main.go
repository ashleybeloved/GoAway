package main

import (
	"goaway/internal/handlers"
	"goaway/internal/repositories"
	"log"
	"os"

	"github.com/gin-contrib/cors"
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

	// CORS

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ADDRESS")},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// Handlers

	r.POST("/reg", handlers.Reg)
	r.POST("/login", handlers.Login)

	profile := r.Group("/u")
	profile.Use(handlers.AuthMiddleware())
	{
		profile.POST("/new", handlers.New)
		profile.POST("/logout", handlers.Logout)
		profile.GET("/links", handlers.Links)
		profile.GET("/links/:id", handlers.Link)
		profile.DELETE("/links/:id", handlers.DelLink)
	}

	r.GET("/:id", handlers.Redirect)

	log.Println("The server is running on :8080")
	r.Run(":" + os.Getenv("PORT"))
}

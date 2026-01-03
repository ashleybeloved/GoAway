package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env

	err := godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatal("Failed to load .env:", err)
	}

	// Set Release Mode for Gin

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Handlers

	r.POST("/register")
	r.POST("/login")
	r.POST("/logout")

	r.POST("/new")
	r.GET("/:id")

	// Run server

	log.Println("The server is running on :8080")
	r.Run(":%v", os.Getenv("PORT"))
}

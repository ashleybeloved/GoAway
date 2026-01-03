package handlers

import (
	"goaway/internal/models"
	"goaway/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := services.Reg(req.Login, req.Password)
	if err != nil {
		switch err {

		}
	}
}

func Login(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := services.Reg(req.Login, req.Password)
	if err != nil {
		switch err {

		}
	}
}

func Logout(c *gin.Context) {

}

func New(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := services.Reg(req.Login, req.Password)
	if err != nil {
		switch err {

		}
	}
}

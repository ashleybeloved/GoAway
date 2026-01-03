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
		case services.ErrInvalidChars:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		case services.ErrLoginTooShort:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		case services.ErrPasswordTooShort:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		case services.ErrUserExists:
			c.JSON(http.StatusConflict, gin.H{"error": err})
			return
		case services.ErrHashPassword:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		case services.ErrCreateUser:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "You are successfully registered!"})
}

func Login(c *gin.Context) {
	var req models.AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err := services.Login(req.Login, req.Password)
	if err != nil {
		switch err {
		case services.ErrUserNotExists:
			c.JSON(http.StatusForbidden, gin.H{"error": err})
			return
		case services.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "You are logged in " + req.Login})
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

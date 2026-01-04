package handlers

import (
	"goaway/internal/models"
	"goaway/internal/services"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func New(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var req models.LinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	shortUrl, err := services.New(req.URL, userID)
	if err != nil {
		switch err {
		case services.ErrInvalidRequest:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		case services.ErrNotLink:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		case services.ErrCreateLink:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Your link: " + os.Getenv("DOMAIN") + "/" + shortUrl})
}

func Redirect(c *gin.Context) {
	shortUrl := c.Param("id")

	URL, err := services.Redirect(shortUrl)
	if err != nil {
		switch err {
		case services.ErrURLNotExists:
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
	}

	c.Redirect(http.StatusFound, URL)
}

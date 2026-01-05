package handlers

import (
	"goaway/internal/models"
	"goaway/internal/services"
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{"message": "create success",
		"link": "http://localhost:8080" + "/" + shortUrl})
}

func Redirect(c *gin.Context) {
	shortUrl := c.Param("id")

	URL, err := services.Redirect(shortUrl)
	if err != nil {
		switch err {
		case services.ErrURLNotExists:
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.Redirect(http.StatusFound, URL)
}

func Links(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	links, err := services.Links(userID)
	if err != nil {
		switch err {
		case services.ErrGetLinks:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid JSON"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, links)
}

func Link(c *gin.Context) {
	shortUrl := c.Param("id")
	userID := c.MustGet("user_id").(uint)

	link, err := services.Link(shortUrl, userID)
	if err != nil {
		switch err {
		case services.ErrLinkNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.JSON(http.StatusOK, link)
}

func DelLink(c *gin.Context) {
	shortUrl := c.Param("id")
	userID := c.MustGet("user_id").(uint)

	err := services.DelLink(shortUrl, userID)
	if err != nil {
		switch err {
		case services.ErrLinkNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "invalid JSON"})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}
}

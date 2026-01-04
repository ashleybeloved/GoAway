package handlers

import (
	"goaway/internal/models"
	"goaway/internal/services"
	"net/http"
	"os"
	"strconv"

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

	sessionToken, err := services.Login(req.Login, req.Password)
	if err != nil {
		switch err {
		case services.ErrUserNotExists:
			c.JSON(http.StatusForbidden, gin.H{"error": err})
			return
		case services.ErrInvalidPassword:
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	timeToLiveStr := os.Getenv("TIME_TO_LIVE")
	timeToLive, _ := strconv.Atoi(timeToLiveStr)

	c.SetCookie(
		"session_token",     // Name of the cookie
		sessionToken,        // Value (the generated UUID)
		3600*timeToLive,     // MaxAge in seconds
		"/",                 // Path (accessible on all routes)
		os.Getenv("DOMAIN"), // Domain
		false,               // Secure: set to true only if using HTTPS
		true,                // HttpOnly: prevents JavaScript from accessing the cookie
	)

	c.JSON(http.StatusOK, gin.H{"message": "You are logged in " + req.Login})
}

func Logout(c *gin.Context) {
	token, err := c.Cookie("session_token")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Already logged out"})
		return
	}

	err = services.Logout(token)
	if err != nil {
		switch err {
		case services.ErrDelSession:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	}

	c.SetCookie(
		"session_token",
		"",
		-1, // MaxAge -1 tells the browser to delete the cookie immediately
		"/",
		os.Getenv("DOMAIN"),
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

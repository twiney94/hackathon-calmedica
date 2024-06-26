package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

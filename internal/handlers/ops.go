package handlers

import (
	"github.com/gin-gonic/gin"
)

// GET /ops/ping
// Get Ping
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

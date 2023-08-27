package handlers

import (
	"github.com/gin-gonic/gin"
)

// Ping godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags ops
// @Accept json
// @Produce json
// @Success 200 {string} pong
// @Router /ops/ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

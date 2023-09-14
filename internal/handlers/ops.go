package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/core/domain"
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

// Info micro service
// @Summary Info
// @Schemes
// @Description Informations sur le service
// @Tags ops
// @Accept json
// @Produce json
// @Success 200 {object} domain.Info
// @Router /ops/info [get]
func Info(c *gin.Context) {
	c.JSON(200, domain.Info{
		Version:     "0.1.0",
		Name:        "ms-mongodb-api",
		Description: "Service d'accès à une base mongodb. Avec une approche générique.",
	})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index Index
func Index(c *gin.Context) {
	c.String(http.StatusOK, "coupon-server index")
}

// Health Health
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "UP"})
}

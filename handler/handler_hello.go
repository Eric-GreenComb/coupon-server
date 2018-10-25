package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetHelloJWT GetHelloJWT
func GetHelloJWT(c *gin.Context) {
	_userID := c.MustGet("userID")
	c.JSON(http.StatusOK, gin.H{
		"sub":  _userID,
		"text": "Get Hello",
	})
}

// PostHelloJWT PostHelloJWT
func PostHelloJWT(c *gin.Context) {
	_userID := c.MustGet("userID")
	c.JSON(http.StatusOK, gin.H{
		"sub":  _userID,
		"text": "Post Hello",
	})
}

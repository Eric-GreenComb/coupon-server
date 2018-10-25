package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/coupon-server/bean"
)

// GetHelloAPI GetHelloAPI
func GetHelloAPI(c *gin.Context) {

	var _formParams bean.FormParams
	err := c.BindJSON(&_formParams)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   _formParams.ID,
		"name": _formParams.Name,
	})
}

// PostHelloAPI PostHelloAPI
func PostHelloAPI(c *gin.Context) {

	var _formParams bean.FormParams
	err := c.BindJSON(&_formParams)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":   _formParams.ID,
		"name": _formParams.Name,
	})
}

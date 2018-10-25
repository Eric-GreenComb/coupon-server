package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	api "github.com/Eric-GreenComb/contrib/wechat"
	"github.com/Eric-GreenComb/coupon-server/config"
	"github.com/gin-gonic/gin"
)

// APIAuth APIAuth
func APIAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 把request的内容读取出来
		var _bodyBytes []byte
		if c.Request.Body == nil {
			APIAbortWithError(c, http.StatusInternalServerError, "Request.Body is Nil")
			return
		}

		_bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		_reader := bytes.NewReader(_bodyBytes)
		var _props map[string]interface{}
		err := api.BindJSON(_reader, &_props)
		if err != nil {
			APIAbortWithError(c, http.StatusInternalServerError, "BindJSON Error")
			return
		}

		_sign := strings.ToLower(api.WxpayCalcSign(_props, config.Server.Key))
		fmt.Println("====== api signed : ", _sign)
		if _props["signed"] != _sign {
			APIAbortWithError(c, http.StatusInternalServerError, "Invaild Signed")
			return
		}
		// 把刚刚读出来的再写进去
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(_bodyBytes))
		c.Next()
	}
}

// APIAbortWithError APIAbortWithError
func APIAbortWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

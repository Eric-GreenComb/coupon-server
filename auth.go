package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/coupon-server/bean"
	"github.com/Eric-GreenComb/coupon-server/config"
	"github.com/Eric-GreenComb/coupon-server/persist"
)

// Login Login
func Login(c *gin.Context) {

	var _userFarams bean.UserParams

	if c.BindJSON(&_userFarams) != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "bind json error."})
		return
	}

	sum := sha256.Sum256([]byte(_userFarams.Pwd))
	_pwd := fmt.Sprintf("%x", sum)

	user, err := persist.GetPersist().Login(_userFarams.UserID, _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "userid or pwd error."})
		return
	}

	expire := time.Now().Add(config.JWT.ExpireTime)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["sub"] = user.UserID
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(config.JWT.SigningKey))
	if err != nil {
		JWTAbortWithError(c, http.StatusUnauthorized, "Create JWT Token faild", config.JWT.Realm)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

// AdminLogin AdminLogin
func AdminLogin(c *gin.Context) {

	var _userFarams bean.UserParams

	if c.BindJSON(&_userFarams) != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "bind json error."})
		return
	}

	sum := sha256.Sum256([]byte(_userFarams.Pwd))
	_pwd := fmt.Sprintf("%x", sum)

	user, err := persist.GetPersist().AdminLogin(_userFarams.UserID, _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "user or pwd error."})
		return
	}

	expire := time.Now().Add(config.JWT.ExpireTime)
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["sub"] = user.UserID
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(config.JWT.SigningKey))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "Create JWT Token faild."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"token":   tokenString,
		"expire":  expire.Format(time.RFC3339),
	})
}

// JWTAuth JWTAuth
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		_token, err := request.ParseFromRequest(c.Request, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(config.JWT.SigningKey))
			return b, nil
		})

		if err != nil {
			JWTAbortWithError(c, http.StatusUnauthorized, "Invaild User Token", config.JWT.Realm)
			return
		}

		claims := _token.Claims.(jwt.MapClaims)

		c.Set("userID", claims["sub"])

		c.Next()
	}
}

// RefreshToken RefreshToken
func RefreshToken(c *gin.Context) {

	var _userFarams bean.UserParams
	if c.BindJSON(&_userFarams) != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "bind json error."})
		return
	}

	_userID := _userFarams.UserID
	expire := time.Now().Add(config.JWT.ExpireTime)

	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["sub"] = _userID
	claims["exp"] = expire.Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	// Set some claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(config.JWT.SigningKey))

	if err != nil {
		JWTAbortWithError(c, http.StatusUnauthorized, "Create JWT Token faild", config.JWT.Realm)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  tokenString,
		"expire": expire.Format(time.RFC3339),
	})
}

// JWTAbortWithError JWTAbortWithError
func JWTAbortWithError(c *gin.Context, code int, message, realm string) {
	c.Header("WWW-Authenticate", "JWT realm="+realm)
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
	c.Abort()
}

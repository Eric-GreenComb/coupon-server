package main

import (
	"log"
	"net/http"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/coupon-server/config"
	"github.com/Eric-GreenComb/coupon-server/handler"
	"github.com/Eric-GreenComb/coupon-server/persist"
)

var (
	g errgroup.Group
)

func main() {
	if config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDatabase()

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)

		r0.POST("login", Login)

		r0.POST("admin/login", AdminLogin)
	}

	// api
	r1 := router.Group("/api/v1")
	{
		// 用户
		r1.POST("/user/create", handler.CreateUser)
		r1.GET("/user/:userid", handler.UserInfo)
		r1.POST("/user/updatepwd", handler.UpdateUserPasswd)
		r1.GET("/users/list/:page/:search", handler.ListUser)
		r1.GET("/users/count", handler.CountUser)
	}

	// admin api
	r2 := router.Group("/api/admin/v1")
	{
		r2.POST("/user/create", handler.CreateAdminUser)
		r2.GET("/user/:userid", handler.AdminUserInfo)
		r2.POST("/user/updatepwd", handler.UpdateAdminUserPasswd)
	}

	// auth api
	r100 := router.Group("/api/auth/v1")
	r100.Use(JWTAuth())
	{
		r100.GET("/hello", handler.GetHelloJWT)
		r100.POST("/hello", handler.PostHelloJWT)
		r100.GET("/refresh_token", RefreshToken)
	}

	// wechat api signed
	r101 := router.Group("/api/wechat/v1")
	r101.Use(APIAuth())
	{
		r101.GET("/hello", handler.GetHelloAPI)
		r101.POST("/hello", handler.PostHelloAPI)
	}

	for _, _port := range config.Server.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return gracehttp.Serve(server)
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

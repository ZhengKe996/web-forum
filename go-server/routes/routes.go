package routes

import (
	"github.com/gin-gonic/gin"
	"go-web-example/controller"
	"go-web-example/logger"
	"go-web-example/middleware"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用中间件

	r.POST("/sign", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		c.String(http.StatusOK, "Pong")
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})
	return r
}

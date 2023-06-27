package routes

import (
	"github.com/gin-gonic/gin"
	"go-web-example/controller"
	"go-web-example/logger"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用中间件

	r.POST("/sign", controller.SignUpHandler)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "404",
		})
	})
	return r
}

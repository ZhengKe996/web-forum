package routes

import (
	"github.com/gin-gonic/gin"
	"go-web-example/logger"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true)) // 使用中间件

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}

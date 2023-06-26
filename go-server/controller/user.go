package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-example/logic"
	"go-web-example/models"
	"go.uber.org/zap"
	"net/http"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUpHandler with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	// 2.业务处理
	logic.SignUp(p)

	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"success": true,
	})
}

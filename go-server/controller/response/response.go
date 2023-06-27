package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseDate struct {
	Code    ResCode     `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

// ResponseError 定义通用错误返回结构
func ResponseError(c *gin.Context, code ResCode) {

	c.JSON(http.StatusOK, &ResponseDate{
		Code:    code,
		Message: code.Message(),
		Data:    nil,
	})
}

// ResponseErrorWithMessage 定义自定义错误返回结构
func ResponseErrorWithMessage(c *gin.Context, code ResCode, message interface{}) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// ResponseSuccess 定义通用正确返回结构
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &ResponseDate{
		Code:    CodeSuccess,
		Message: CodeSuccess.Message(),
		Data:    data,
	})
}

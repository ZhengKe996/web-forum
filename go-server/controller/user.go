package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-web-example/controller/response"
	"go-web-example/dao/mysql"
	"go-web-example/logic"
	"go-web-example/models"
	"go.uber.org/zap"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数与参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUpHandler with invalid param", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("SignUpHandler insert error", zap.Error(err))

		if errors.Is(err, mysql.ErrorUserExist) {
			response.ResponseError(c, response.CodeUserExist)
			return
		}
		fmt.Println(err.Error())

		response.ResponseErrorWithMessage(c, response.CodeServerBusy, err.Error())
		return
	}

	// 3.返回响应
	response.ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	// 1.获取参数与参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("LoginHandler with invalid param", zap.Error(err))
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}

	// 2.业务处理
	if token, err := logic.Login(p); err != nil {
		zap.L().Error("LoginHandler query error", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			response.ResponseError(c, response.CodeUserNotExist)
			return
		}

		response.ResponseErrorWithMessage(c, response.CodeInvalidParam, err.Error())
		return
	} else {
		// 3.返回响应
		response.ResponseSuccess(c, token)
	}
}

package logic

import (
	"go-web-example/dao/mysql/user"
	"go-web-example/models"
	"go-web-example/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 1. 判断用户是否存在
	sqlUser.QueryUserByUsername()

	// 2. 生成UID
	snowflake.GenID()
	// 3. 保存进数据库
	sqlUser.InsertUser()
}

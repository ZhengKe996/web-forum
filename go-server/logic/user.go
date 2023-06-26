package logic

import (
	"go-web-example/dao/mysql"
	"go-web-example/models"
	"go-web-example/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()

	// 3.构造一个User实例

	u := &models.User{UserID: userID, Username: p.Username, Password: p.Password}

	// 4.保存进数据库
	return mysql.InsertUser(u)
}

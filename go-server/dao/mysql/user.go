package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go-web-example/models"
)

const secret = "HelloWord"

// InsertUser 插入一条新的用户记录
func InsertUser(u *models.User) error {
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	// 对密码进行加密
	_, err := db.Exec(sqlStr, u.UserID, u.Username, encryptPassword(u.Password))
	if err != nil {
		return err
	}
	return nil
}

// CheckUserExist 依据用户名查询
func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return nil
}

// encryptPassword 对密码进行加密
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

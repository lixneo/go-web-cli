package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go-web-cli/models"
	"go-web-cli/pkg/snowflake"
)

const secret = "soSerious"

func encryptPassword(data []byte) (result string) {
	h := md5.New()
	h.Write([]byte(secret))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func Register(user *models.User) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	err = db.Get(&count, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if count > 0 {
		// 用户已存在
		return errors.New("user exists")
	}
	// 生成user_id
	userID, err := snowflake.GetID()
	if err != nil {
		return errors.New("get user id err:" + err.Error())
	}
	// 生成加密密码
	password := encryptPassword([]byte(user.Password))
	// 把用户插入数据库
	sqlStr = "insert into user(user_id, username, password) values (?,?,?)"
	_, err = db.Exec(sqlStr, userID, user.UserName, password)
	return
}

func Login(user *models.User) (err error) {
	originPassword := user.Password // 记录一下原始密码
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.UserName)
	if err != nil && err != sql.ErrNoRows {
		// 查询数据库出错
		return
	}
	if err == sql.ErrNoRows {
		// 用户不存在
		return ErrorUserNotExit
	}
	// 生成加密密码与查询到的密码比较
	password := encryptPassword([]byte(originPassword))
	if user.Password != password {
		return ErrorPasswordWrong
	}
	return
}

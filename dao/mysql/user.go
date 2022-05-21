package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"webapp-scaffold/models"
)

const secret = "2zyyyyy"

var (
	ErrorUserExist       = errors.New("该用户已存在")
	ErrorUserNotExists   = errors.New("该用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// CheckUserExist 检查用户名是否存在
func CheckUserExist(username string) (err error) {
	var count int
	sqlStr := "select count(user_id) from user where username = ?"
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser 写入用户数据
func InsertUser(u *models.User) (err error) {
	// 密码加密
	u.Password = encryptPassword(u.Password)
	//执行sql语句写入数据
	sqlStr := "insert into user(user_id, username, password, email, gender) value(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, u.UserID, u.UserName, u.Password, u.Email, u.Gender)
	return
}

// encryptPassword 密码加密
func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

// Login 用户登录
func Login(u *models.User) (err error) {
	oPassword := u.Password // 用户登录的密码
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(u, sqlStr, u.UserName)
	if err == sql.ErrNoRows {
		// 如果没有查询到该用户
		return ErrorUserNotExists
	}
	if err != nil {
		// 数据库查询失败
		return err
	}
	// 判断密码是否正确
	password := encryptPassword(oPassword)
	if password != u.Password {
		return ErrorInvalidPassword
	}
	return
}

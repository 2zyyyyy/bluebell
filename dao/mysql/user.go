package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"webapp-scaffold/models"
)

const secret = "2zyyyyy"

// CheckUserExist 检查用户名是否存在
func CheckUserExist(username string) (err error) {
	var count int
	sql := "select count(user_id) from user where username = ?"
	if err := db.Get(&count, sql, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("该用户已存在")
	}
	return
}

// InsertUser 写入用户数据
func InsertUser(u *models.User) (err error) {
	// 密码加密
	u.Password = encryptPassword(u.Password)
	//执行sql语句写入数据
	sql := "insert into user(user_id, username, password, email, gender) value(?,?,?,?,?)"
	_, err = db.Exec(sql, u.UserID, u.UserName, u.Password, u.Email, u.Gender)
	return
}

func encryptPassword(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

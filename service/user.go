package service

import (
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/models"
	"webapp-scaffold/pkg/jwt"
	"webapp-scaffold/pkg/snowflake"
)

// SignUp 处理用户注册逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2.生成UID
	userID, err := snowflake.GetID()
	u := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
		Email:    p.Email,
		Gender:   p.Gender,
	}
	// 3.保存用户数据到user表
	return mysql.InsertUser(u)
}

// Login 处理用户登录逻辑
func Login(p *models.ParamLogin) (token string, err error) {
	u := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(u); err != nil {
		return "", err
	}
	// 生成JWT
	return jwt.GenToken(u.UserID, u.UserName)
}

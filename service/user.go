package service

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

// 用户相关

// SignUp 处理用户注册逻辑
func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错
		return err
	}
	// 2.生成UID
	userID, err := snowflake.GenID()
	if err != nil {
		zap.L().Error("user snowflake.GenID failed.", zap.Error(err))
		return
	}
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
func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		UserName: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	user.Token, err = jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	return
}

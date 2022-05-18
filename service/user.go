package service

import (
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/models"
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

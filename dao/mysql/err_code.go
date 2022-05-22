package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("该用户已存在")
	ErrorUserNotExists   = errors.New("该用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)

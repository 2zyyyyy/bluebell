package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID "

var ErrorUserNotLogin = errors.New("当前用户未登录")

func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

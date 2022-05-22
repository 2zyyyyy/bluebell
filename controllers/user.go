package controllers

import (
	"errors"
	"fmt"
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/models"
	"webapp-scaffold/service"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// 用户相关

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Println(*p)
	// 2.业务处理
	if err := service.SignUp(p); err != nil {
		zap.L().Error("service.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			// 如果用户已存在
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	// 1.获取请求参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 具体参数错误
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	fmt.Println(*p)
	// 2.业务逻辑处理
	token, err := service.Login(p)
	if err != nil {
		zap.L().Error("service.Login failed.", zap.String("用户:", p.Username), zap.Error(err))
		// 判断用户是否不存在
		if errors.Is(err, mysql.ErrorUserNotExists) {
			ResponseError(c, CodeUserNoExist)
			return
		}
		// 处理其他错误
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 返回响应
	ResponseSuccess(c, token)
}

package controllers

import (
	"fmt"
	"net/http"
	"webapp-scaffold/models"
	"webapp-scaffold/service"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误信息
		})
		return
	}
	fmt.Println(*p)
	// 2.业务处理
	if err := service.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	// 3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误信息
		})
		return
	}
	fmt.Println(*p)
	// 2.业务逻辑处理
	if err := service.Login(p); err != nil {
		zap.L().Error("service.Login failed.", zap.String("用户:", p.Username), zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录失败，用户名或密码错误",
		})
		return
	}
	// 返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

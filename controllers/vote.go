package controllers

import (
	"webapp-scaffold/models"
	"webapp-scaffold/service"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

// 帖子投票

func CommunityVote(c *gin.Context) {
	// 参数校验
	p := new(models.ParamCommunityVote)
	if err := c.ShouldBindJSON(p); err != nil {
		// 类型断言
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	// 获取用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := service.CommunityVote(userID, p); err != nil {
		zap.L().Error(" service.CommunityVote failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

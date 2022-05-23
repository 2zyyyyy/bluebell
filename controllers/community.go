package controllers

import (
	"strconv"
	"webapp-scaffold/models"
	"webapp-scaffold/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// 社区相关

// CommunityHandler 处理获取社区列表的函数
func CommunityHandler(c *gin.Context) {
	// 1.查询到所有社区的信息(community_id, community_name)
	list, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, list)
}

// CommunityDetailHandler 处理获取社区详情的函数
func CommunityDetailHandler(c *gin.Context) {
	// 1.拿到id
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 根据社区id查询社区详情
	detail, err := service.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("service.GetCommunityList failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, detail)
}

// CreatePostHandler 创建帖子函数
func CreatePostHandler(c *gin.Context) {
	// 1.获取参数
	post := new(models.CommunityPost)
	if err := c.ShouldBindJSON(post); err != nil {
		// 如果参数异常就记录日志并返回错误
		zap.L().Debug("c.ShouldBindJSON(post) failed.", zap.Any("err", err))
		zap.L().Error("create community failed.", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.获取用户id
	userId, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	post.AuthorID = int64(userId)
	// 3.参数校验

	// 4.存储数据
	if err := service.CreateCommunityPost(post); err != nil {
		// 创建失败 返回错误信息
		zap.L().Error("service.CreateCommunityPost failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 4.返回响应
	ResponseSuccess(c, post)
}

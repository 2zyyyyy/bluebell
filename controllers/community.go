package controllers

import (
	"strconv"
	"webapp-scaffold/models"
	"webapp-scaffold/service"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const (
	page       = 10
	size       = 1
	orderTime  = "time"
	orderScore = "score"
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
	// 1.获取参数并校验
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
	// 3.存储数据
	if err := service.CreateCommunityPost(post); err != nil {
		// 创建失败 返回错误信息
		zap.L().Error("service.CreateCommunityPost failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 4.返回响应
	ResponseSuccess(c, nil)
}

// PostDetailHandler 帖子详情函数
func PostDetailHandler(c *gin.Context) {
	// 1.拿到postId
	id := c.Param("id")
	postId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		zap.L().Error("get post detail failed. invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	// 根据帖子id查询帖子详情
	detail, err := service.GetPostDetail(uint64(postId))
	if err != nil {
		zap.L().Error("service.GetPostDetail failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, detail)
}

// GetPostListHandler 获取帖子列表函数
func GetPostListHandler(c *gin.Context) {
	page, size, err := getPageInfo(c)
	if err != nil {
		page = 1
		size = 10
	}
	// 2.获取数据
	list, err := service.GetPostList(page, size)
	if err != nil {
		zap.L().Error("service.GetPostList failed.", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回响应
	ResponseSuccess(c, list)
}

// GetPostOrderListHandler 根据指定的排序方式返回数据
func GetPostOrderListHandler(c *gin.Context) {
	// 初始化结构体并指定默认参数值
	p := &models.ParamOrderList{
		Page:  page,
		Size:  size,
		Order: orderTime,
	}
	// 1.获取参数
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.去redis查询id列表

	// 3.根据id去数据库查询帖子详情信息

	ResponseSuccess(c, nil)
	return
}

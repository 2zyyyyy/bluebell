package routers

import (
	"webapp-scaffold/controllers"
	"webapp-scaffold/logger"
	"webapp-scaffold/middlewares"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.New()
	// 使用自定义的中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 初始化翻译器
	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Error("controllers.InitTrans", zap.Error(err))
		return nil
	}

	v1 := r.Group("/api/v1")

	// 用户注册
	v1.POST("/signup", controllers.SignUpHandler)
	// 用户登录
	v1.POST("/login", controllers.LoginHandler)

	// JWT认证
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)           // 社区列表
		v1.GET("/community/:id", controllers.CommunityDetailHandler) // 社区详情

		v1.POST("/community/post", controllers.CreatePostHandler)    // 创建帖子
		v1.GET("/community/post/:id", controllers.PostDetailHandler) // 帖子详情
	}

	return r
}

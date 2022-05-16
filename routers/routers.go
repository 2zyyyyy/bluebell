package routers

import (
	"net/http"
	"webapp-scaffold/controllers"
	"webapp-scaffold/logger"

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

	// 用户注册
	r.POST("/signup", controllers.SignUpHandle)

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	return r
}

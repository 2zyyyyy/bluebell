package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp-scaffold/dao/mysql"
	"webapp-scaffold/dao/redis"
	"webapp-scaffold/logger"
	"webapp-scaffold/pkg/snowflake"
	"webapp-scaffold/routers"
	"webapp-scaffold/settings"

	"go.uber.org/zap"
)

func main() {
	// 1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Println("init setting failed, err:", err)
	}
	// 2.初始化日志
	if err := logger.Init(settings.Config.LogConfig, settings.Config.Mode); err != nil {
		fmt.Println("init logger failed, err:", err)
	}
	// 延迟注册zap
	defer func(l *zap.Logger) {
		_ = l.Sync()
	}(zap.L())
	zap.L().Debug("logger init success!!!")
	// 3.初始化mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Println("init mysql failed, err:", err)
	}
	defer mysql.Close()
	// 雪花算法生成用户id
	if err := snowflake.Init(settings.Config.StartTime, settings.Config.MachineId); err != nil {
		zap.L().Error("init snowflake failed.", zap.Error(err))
		return
	}
	// 4.初始化redis连接
	if err := redis.Init(); err != nil {
		fmt.Println("init redis failed, err:", err)
	}
	defer redis.Close()

	// 5.注册路由
	r := routers.SetUpRouter()
	// 6.启动服务（优雅关机）
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Config.Port),
		Handler: r,
	}
	zap.L().Info(server.Addr)

	// 开启一个goroutine启动服务
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			zap.L().Error("listen failed", zap.Error(err))
		}
	}()
	// 等待中断信号来关闭服务器 为关闭服务器设置一个5s的超市
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	/*
		kill 默认会发送syscall.SIGTERM 信号
		kill -2发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
		kill -9发送 syscall.SIGKILL信号，但是不能被捕获，所以不需要添加它
		signal.Notify把收到的syscall.SIGINT或syscall.SIGTERM信号转发给qui
		t*/
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此 当接收到上述两种信号的时候才会往下执行
	zap.L().Info("shutdown server...")
	// 创建一个5s超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5s内关闭服务（处理未完成的请求后在关闭服务） 超过5s就超时退出
	if err := server.Shutdown(ctx); err != nil {
		zap.L().Error("server shutdown failed.", zap.Error(err))
	}
	zap.L().Info("server exiting")
}

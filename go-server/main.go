package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go-web-example/dao/mysql"
	"go-web-example/dao/redis"
	"go-web-example/logger"
	"go-web-example/routes"
	"go-web-example/settings"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
	GoWeb 较通用脚手架模版
*/

func main() {
	var filePath string
	flag.StringVar(&filePath, "config", "./config.yaml", "配置文件")
	flag.Parse()

	// 1. 加载配置文件（本地）
	if err := settings.Init(filePath); err != nil {
		fmt.Printf("Init settings failed,err:%v\n", err)
	}
	// 2. 初始化日志
	if err := logger.Init(settings.Conf.LogConfig); err != nil {
		fmt.Printf("Init logger failed,err:%v\n", err)
	}
	defer zap.L().Sync() // 延迟注册
	zap.L().Debug("Init logger success~")

	// 3. 初始化数据库连接
	if err := mysql.InitDB(settings.Conf.MYSQLConfig); err != nil {
		zap.L().Error("Init mysql connect failed", zap.Error(err))
	}
	defer mysql.Close()

	if err := redis.InitDB(settings.Conf.RedisConfig); err != nil {
		zap.L().Error("Init redis connect failed", zap.Error(err))
	}
	defer redis.Close()

	// 4. 注册路由
	router := routes.Setup()
	// 5. 启动服务（优雅关机）
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", viper.GetString("app.port")),
		Handler: router,
	}

	go func() {
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}

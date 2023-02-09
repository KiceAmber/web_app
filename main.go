package main

import (
	"fmt"
	"os"
	"web_app/config"
	"web_app/dao/mysql"
	"web_app/log"
	"web_app/router"

	"go.uber.org/zap"
)

// Go Web 开发通用脚手架

func main() {
	if len(os.Args) < 2 {
		fmt.Print("Need configig File arguments...")
		return
	}
	// 1. 加载配置文件
	if err := config.Init(os.Args[1]); err != nil {
		fmt.Printf("init config failed, err:%v\n", err)
		return
	}

	// 2. 初始化日志
	if err := log.Init(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success...")

	// 3. 初始化 MySQL 连接
	if err := mysql.Init(config.Conf.MysqlConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化 Redis 连接
	// if err := redis.Init(config.Conf.RedisConfig); err != nil {
	// 	fmt.Printf("init redis failed, err:%v\n", err)
	// 	return
	// }
	// defer redis.Close()

	// 5. 注册路由 router
	r := router.Init()

	// 6. 启动服务(优雅关机)
	router.Setup(r)
}

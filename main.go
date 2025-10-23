package main

import (
	"venturan/bootstrap"
	"venturan/global"
)

// @title 基于Golang Gin框架的EAMS系统
// @version 1.0
// @description 基于Golang Gin框架的EAMS系统
// @termsofservice https://github.com/18211167516/Go-Gin-Api
// @contact.name sutter
// @contact.email wu_yiqiang@outlook.com
// @host 127.0.0.1:9527
// @securityDefinitions.apikey Auth
// @in header
// @name Authorization
func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 初始化验证器
	bootstrap.InitializeValidator()
	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()
	// 初始化定时任务
	bootstrap.InitializeCron()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()
	// 启动服务
	bootstrap.RunServer()
}

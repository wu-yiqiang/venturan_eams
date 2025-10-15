package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"venturan/bootstrap"
	"venturan/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()
	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	// 启动服务器
	r.Run(":" + global.App.Config.App.Port)
}

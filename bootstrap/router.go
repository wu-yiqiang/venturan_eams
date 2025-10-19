package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"venturan/app/middleware"
	"venturan/docs"
	"venturan/global"
	"venturan/routes"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	// 跨域
	router.Use(middleware.Cors())
	// swagger
	docs.SwaggerInfo.Title = global.App.Config.Swagger.Title
	docs.SwaggerInfo.Description = global.App.Config.Swagger.Desc
	docs.SwaggerInfo.Host = global.App.Config.Swagger.Host + ":" + global.App.Config.App.Port
	docs.SwaggerInfo.BasePath = global.App.Config.Swagger.BasePath
	docs.SwaggerInfo.Version = global.App.Config.Swagger.Version
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.Static("/html", "./public")
	// 设置 swagger访问路由
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 前端项目静态资源
	router.StaticFile("/", "./static/dist/index.html")
	router.Static("/assets", "./static/dist/assets")
	router.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	router.Static("/public", "./static")
	router.Static("/storage", "./storage/app/public")

	// 注册 api 分组路由
	apiGroup := router.Group("/")
	routes.SetApiGroupRoutes(apiGroup)
	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

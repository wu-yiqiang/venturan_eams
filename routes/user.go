package routes

import (
	"github.com/gin-gonic/gin"
	"venturan/app/controllers"
	"venturan/app/middleware"
	"venturan/app/services"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.DELETE("/user/:user_id", controllers.Delete)
		authRouter.POST("/user/info", controllers.Info)
		authRouter.POST("/user/logout", controllers.Logout)
	}
}
